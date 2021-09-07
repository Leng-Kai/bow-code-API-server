package classroom

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	// "strings"
	"time"

	. "github.com/Leng-Kai/bow-code-API-server/course_plan"
	"github.com/Leng-Kai/bow-code-API-server/db"
	"github.com/Leng-Kai/bow-code-API-server/schemas"
	"github.com/Leng-Kai/bow-code-API-server/user"
	"github.com/Leng-Kai/bow-code-API-server/util"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {

}

func GetClassrooms(w http.ResponseWriter, r *http.Request) {

}

func CreateNewClassroom(w http.ResponseWriter, r *http.Request) {
	body, err := util.GetBody(r)
	if err != nil {
		// http.Error()
		return
	}
	newClassroom := schemas.Classroom{}
	err = json.Unmarshal(body, &newClassroom)
	if err != nil {
		// http.Error()
		return
	}
	creator, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	newClassroom.Creator = creator.UserID
	newClassroom.CreateTime = time.Now()
	newClassroom.Students = []schemas.UserID{}
	newClassroom.Applicants = []schemas.UserID{}
	newClassroom.Invitees = []schemas.UserID{}
	newClassroom.CourseList = []schemas.ClassroomComponent{}
	newClassroom.HomeworkList = []schemas.ClassroomComponent{}
	newClassroom.ExamList = []schemas.ClassroomComponent{}

	coursePlan, _ := db.GetSingleCoursePlan(bson.D{{"_id", newClassroom.CoursePlan}}, bson.D{})
	for _, component := range coursePlan.ComponentList {
		classroomComponent := schemas.ClassroomComponent{ Component: component, Begin: -1, End: -1, Private: false }
		if component.Type == COURSE {
			newClassroom.CourseList = append(newClassroom.CourseList, classroomComponent)
		} else if component.Type == HOMEWORK {
			newClassroom.HomeworkList = append(newClassroom.HomeworkList, classroomComponent)
		} else if component.Type == EXAM {
			newClassroom.ExamList = append(newClassroom.ExamList, classroomComponent)
		}
	}

	id, err := db.CreateClassroom(newClassroom)
	if err != nil {
		log.Println(err)
		// http.Error()
		return
	}
	_, err = db.UpdateUser(bson.D{{"_id", creator.UserID}}, bson.D{{"$push", bson.D{{"ownClassroomList", id}}}}, true)
	if err != nil {
		log.Println(err)
		// http.Error()
		return
	}

	newClassroomRecord := schemas.ClassroomRecord{ ClassroomRecordID: id, ScoreEntryList: []schemas.ScoreEntry{} }
	crrid, err := db.CreateClassroomRecord(newClassroomRecord)
	if err != nil {
		log.Println(err)
		// http.Error()
		return
	}
	if crrid != id {
		http.Error(w, "unexpected error.", 404)
		return
	}

	resp := struct {
		ClassroomID schemas.ID
	}{ClassroomID: id}
	util.ResponseJSON(w, resp)
}

func ApplyForClassroom(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["crid"]
	crid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", crid}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "classroom not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	// TODO: Prevent replication
	if util.Contain_str(classroom.Students, uid) {
		http.Error(w, "Already applied.", 404)
		return
	}

	// TODO: Prevent application for owned classroom

	update := bson.D{}

	if !classroom.Apply {
		http.Error(w, "classroom not available.", 404)
		return
	} else if classroom.Review {
		filter = bson.D{{"_id", uid}}
		update = bson.D{{"$push", bson.D{{"appliedClassroomList", crid}}}}
		_, err = db.UpdateUser(filter, update, false)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		filter = bson.D{{"_id", crid}}
		update = bson.D{{"$push", bson.D{{"applicants", uid}}}}
		_, err = db.UpdateClassroom(filter, update, false)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
	} else {
		filter = bson.D{{"_id", uid}}
		update = bson.D{{"$push", bson.D{{"joinedClassroomList", crid}}}}
		_, err = db.UpdateUser(filter, update, false)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		filter = bson.D{{"_id", crid}}
		update = bson.D{{"$push", bson.D{{"students", uid}}}}
		_, err = db.UpdateClassroom(filter, update, false)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		// err = AddRecordsForNewStudent(crid, uid)
		// if err != nil {
		// 	http.Error(w, err.Error(), 404)
		// 	return
		// }
	}

	w.WriteHeader(200)
}

func AcceptApplication(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["crid"]
	crid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", crid}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "classroom not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	if uid != classroom.Creator {
		http.Error(w, "permission denied. not classroom creator.", 401)
		return
	}

	uid = mux.Vars(r)["uid"]
	filter = bson.D{{"_id", uid}}
	sortby = bson.D{}
	user_obj, err = db.GetSingleUser(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "user not found.", 404)
		return
	}

	update := bson.D{}

	if !util.Contain_str(classroom.Applicants, uid) || !util.Contain_ID(user_obj.AppliedClassroomList, crid) {
		http.Error(w, "user had not applied for the classroom.", 404)
		return
	}

	if !util.Contain_ID(user_obj.JoinedClassroomList, crid) {
		filter = bson.D{{"_id", uid}}
		update = bson.D{{"$push", bson.D{{"appliedClassroomList", crid}}}}
		_, err = db.UpdateUser(filter, update, false)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
	}

	if util.Contain_str(classroom.Students, uid) {
		http.Error(w, "user had already been added into the classroom.", 404)
		return
	}

	filter = bson.D{{"_id", crid}}
	update = bson.D{{"$push", bson.D{{"students", uid}}}}
	_, err = db.UpdateClassroom(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	filter = bson.D{{"_id", uid}}
	update = bson.D{{"$pull", bson.D{{"appliedClassroomList", crid}}}}
	_, err = db.UpdateUser(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	filter = bson.D{{"_id", crid}}
	update = bson.D{{"$pull", bson.D{{"applicants", uid}}}}
	_, err = db.UpdateClassroom(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	// err = AddRecordsForNewStudent(crid, uid)
	// if err != nil {
	// 	http.Error(w, err.Error(), 404)
	// 	return
	// }

	w.WriteHeader(200)
}

func InviteStudent(w http.ResponseWriter, r *http.Request) {

}

func JoinClassroom(w http.ResponseWriter, r *http.Request) {

}

func GetClassroomStatus(w http.ResponseWriter, r *http.Request) {

}

func GetClassroomByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["crid"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", objId}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "classroom not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	// if uid != classroom.Creator {
	// 	http.Error(w, "permission denied. not classroom creator.", 401)
	// 	return
	// }

	// modify classroom if corresponding course plan is modified

	resp := struct {
		Classroom schemas.Classroom `json:"classroom"`
		IsCreator bool                `json:"isCreator"`
	}{Classroom: classroom, IsCreator: (uid == classroom.Creator)}
	util.ResponseJSON(w, resp)
}

func UpdateClassroomByID(w http.ResponseWriter, r *http.Request) {
	body, err := util.GetBody(r)
	if err != nil {
		// http.Error()
		return
	}
	id := mux.Vars(r)["crid"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", objId}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "classroom not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	if uid != classroom.Creator {
		http.Error(w, "permission denied. not classroom creator.", 401)
		return
	}

	var updatedClassroom map[string]interface{}
	err = json.Unmarshal(body, &updatedClassroom)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	// prevent other attributes from being changed

	for key, value := range updatedClassroom {
		update := bson.D{{"$set", bson.D{{key, value}}}}
		_, err = db.UpdateClassroom(filter, update, false)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
	}
}

func GetClassroomRecord(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["crid"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", objId}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "classroom not found.", 404)
		return
	}
	classroomRecord, err := db.GetSingleClassroomRecord(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "classroom record not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	if uid != classroom.Creator {
		http.Error(w, "permission denied. not classroom creator.", 401)
		return
	}

	util.ResponseJSON(w, classroomRecord)
}

func GetStudentScores(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["crid"]
	crid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
		return
	}
	uid_student := mux.Vars(r)["uid"]

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	filter := bson.D{{"_id", crid}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		log.Println(err)
		http.Error(w, "classroom not found.", 404)
		return
	}
	if uid != classroom.Creator && uid != uid_student {
		http.Error(w, "permission denied.", 401)
		return
	}
	
	classroomRecord, err := db.GetSingleClassroomRecord(filter, sortby)
	if err != nil {
		log.Println(err)
		http.Error(w, "classroom record not found.", 404)
		return
	}

	scoreEntryList := classroomRecord.ScoreEntryList
	// scoreEntries := []schemas.ScoreEntry{}
	problemScore := map[schemas.ProblemID]int{}
	for _, scoreEntry := range scoreEntryList {
		if scoreEntry.UserID == uid_student {
			// scoreEntries = append(scoreEntries, scoreEntry)
			problemScore[scoreEntry.ProblemID] = scoreEntry.Score
		}
	}

	homeworkList := classroom.HomeworkList
	homeworkComponentScoreList := []schemas.ComponentScore{}
	for _, homework := range homeworkList {
		setScoreList := []schemas.SetScore{}
		for _, set := range homework.Component.SetList {
			score := -1
			if s, exist := problemScore[set.ID]; exist {
				score = s
			}
			setScoreList = append(setScoreList, schemas.SetScore{Name: set.Name, Score: score})
		}
		componentScore := schemas.ComponentScore{Name: homework.Component.Name, SetScoreList: setScoreList}
		homeworkComponentScoreList = append(homeworkComponentScoreList, componentScore)
	}

	examList := classroom.ExamList
	examComponentScoreList := []schemas.ComponentScore{}
	for _, exam := range examList {
		setScoreList := []schemas.SetScore{}
		for _, set := range exam.Component.SetList {
			score := -1
			if s, exist := problemScore[set.ID]; exist {
				score = s
			}
			setScoreList = append(setScoreList, schemas.SetScore{Name: set.Name, Score: score})
		}
		componentScore := schemas.ComponentScore{Name: exam.Component.Name, SetScoreList: setScoreList}
		examComponentScoreList = append(examComponentScoreList, componentScore)
	}

	resp := struct {
		HomeworkComponentScoreList []schemas.ComponentScore `json:"homeworkComponentScoreList"`
		ExamComponentScoreList     []schemas.ComponentScore `json:"examComponentScoreList"`
	}{HomeworkComponentScoreList: homeworkComponentScoreList, ExamComponentScoreList: examComponentScoreList}
	util.ResponseJSON(w, resp)
}

func CreateHomework(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["crid"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", objId}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		log.Println(err)
		http.Error(w, "classroom not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	if uid != classroom.Creator {
		http.Error(w, "permission denied. not classroom creator.", 401)
		return
	}

	body, err := util.GetBody(r)
	if err != nil {
		// http.Error()
		return
	}
	newClassroomComponent := schemas.ClassroomComponent{}
	err = json.Unmarshal(body, &newClassroomComponent)
	if err != nil {
		// http.Error()
		return
	}

	update := bson.D{{"$push", bson.D{{"homeworkList", newClassroomComponent}}}}
	_, err = db.UpdateClassroom(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}

func CreateExam(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["crid"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", objId}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		log.Println(err)
		http.Error(w, "classroom not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	if uid != classroom.Creator {
		http.Error(w, "permission denied. not classroom creator.", 401)
		return
	}

	body, err := util.GetBody(r)
	if err != nil {
		// http.Error()
		return
	}
	newClassroomComponent := schemas.ClassroomComponent{}
	err = json.Unmarshal(body, &newClassroomComponent)
	if err != nil {
		// http.Error()
		return
	}

	update := bson.D{{"$push", bson.D{{"examList", newClassroomComponent}}}}
	_, err = db.UpdateClassroom(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}

func UpdateHomework(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["crid"]
	crid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", crid}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		log.Println(err)
		http.Error(w, "classroom not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID
	if uid != classroom.Creator {
		http.Error(w, "permission denied. not classroom creator.", 401)
		return
	}

	no, err := strconv.Atoi(mux.Vars(r)["No"])
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	body, err := util.GetBody(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	var updatedHomework struct {
		Component schemas.CoursePlanComponent `json:"component"`
		Begin     int  				  		  `json:"begin"`
		End       int  				  		  `json:"end"`
		Private   bool 				  		  `json:"private`
	}
	err = json.Unmarshal(body, &updatedHomework)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	classroom.HomeworkList[no].Component = updatedHomework.Component
	classroom.HomeworkList[no].Begin = updatedHomework.Begin
	classroom.HomeworkList[no].End = updatedHomework.End
	classroom.HomeworkList[no].Private = updatedHomework.Private

	update := bson.D{{"$set", bson.D{{"homeworkList", classroom.HomeworkList}}}}
	_, err = db.UpdateClassroom(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}

func UpdateExam(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["crid"]
	crid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", crid}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		log.Println(err)
		http.Error(w, "classroom not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID
	if uid != classroom.Creator {
		http.Error(w, "permission denied. not classroom creator.", 401)
		return
	}

	no, err := strconv.Atoi(mux.Vars(r)["No"])
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	body, err := util.GetBody(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	var updatedExam struct {
		Component schemas.CoursePlanComponent `json:"component"`
		Begin   int  `json:"begin"`
		End     int  `json:"end"`
		Private bool `json:"private`
	}
	err = json.Unmarshal(body, &updatedExam)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	classroom.ExamList[no].Component = updatedExam.Component
	classroom.ExamList[no].Begin = updatedExam.Begin
	classroom.ExamList[no].End = updatedExam.End
	classroom.ExamList[no].Private = updatedExam.Private

	update := bson.D{{"$set", bson.D{{"examList", classroom.ExamList}}}}
	_, err = db.UpdateClassroom(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}
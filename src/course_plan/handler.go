package course_plan

import (
	"encoding/json"
	"log"
	"net/http"
	// "strconv"
	"strings"
	"time"

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

func GetCoursePlans(w http.ResponseWriter, r *http.Request) {
	
}

func CreateNewCoursePlan(w http.ResponseWriter, r *http.Request) {
	body, err := util.GetBody(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	newCoursePlan := schemas.CoursePlan{}
	err = json.Unmarshal(body, &newCoursePlan)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	creator, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	newCoursePlan.Creator = creator.UserID
	newCoursePlan.CreateTime = time.Now()

	var course schemas.Course
	var problem schemas.Problem

	for i, component := range newCoursePlan.ComponentList {
		if component.Type == COURSE {
			for j, set := range component.SetList {
				course, _ = db.GetSingleCourse(bson.D{{"_id", set.ID}}, bson.D{})
				newCoursePlan.ComponentList[i].SetList[j].Name = course.Name
			}
		} else {
			for j, set := range component.SetList {
				problem, _ = db.GetSingleProblem(bson.D{{"_id", set.ID}}, bson.D{})
				newCoursePlan.ComponentList[i].SetList[j].Name = problem.Name
				newCoursePlan.ComponentList[i].SetList[j].TotalScore = problem.TotalScore
			}
		}
	}

	id, err := db.CreateCoursePlan(newCoursePlan)
	if err != nil {
		log.Println(err)
		// http.Error()
		return
	}
	_, err = db.UpdateUser(bson.D{{"_id", creator.UserID}}, bson.D{{"$push", bson.D{{"ownCoursePlanList", id}}}}, true)
	if err != nil {
		log.Println(err)
		// http.Error()
		return
	}

	resp := struct {
		CoursePlanID schemas.ID
	}{CoursePlanID: id}
	util.ResponseJSON(w, resp)
}

func GetMultipleCoursePlans(w http.ResponseWriter, r *http.Request) {
	coursePlans := r.URL.Query().Get("courseplans")
	coursePlanIDList_str := strings.Split(coursePlans, ",")
	coursePlanIDList := []schemas.CoursePlanID{}
	for _, coursePlanID_str := range coursePlanIDList_str {
		coursePlanID_objID, err := primitive.ObjectIDFromHex(coursePlanID_str)
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}
		coursePlanIDList = append(coursePlanIDList, coursePlanID_objID)
	}

	coursePlanList := []schemas.CoursePlan{}
	for _, cpid := range coursePlanIDList {
		coursePlan, _ := db.GetSingleCoursePlan(bson.D{{"_id", cpid}}, bson.D{})
		coursePlanList = append(coursePlanList, coursePlan)
	}

	resp := struct {
		CoursePlanList []schemas.CoursePlan `json:"coursePlanList"`
	}{CoursePlanList: coursePlanList}
	util.ResponseJSON(w, resp)
}

func GetCoursePlanByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["cpid"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", objId}}
	sortby := bson.D{}
	coursePlan, err := db.GetSingleCoursePlan(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "course plan not found.", 404)
		return
	}
	util.ResponseJSON(w, coursePlan)
}

func UpdateCoursePlanByID(w http.ResponseWriter, r *http.Request) {
	body, err := util.GetBody(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	id := mux.Vars(r)["cpid"]
	cpid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	filter := bson.D{{"_id", cpid}}
	sortby := bson.D{}
	coursePlan, err := db.GetSingleCoursePlan(filter, sortby)
	if err != nil {
		http.Error(w, "course plan not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	if uid != coursePlan.Creator {
		http.Error(w, "permission denied. not course plan creator.", 401)
		return
	}

	var updatedCoursePlan map[string]interface{}
	err = json.Unmarshal(body, &updatedCoursePlan)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	// prevent other attributes from being changed
	legal_key := []string{"name", "componentList", "visibility"}

	for key, value := range updatedCoursePlan {
		if util.Contain_str(legal_key, key) {
			update := bson.D{{"$set", bson.D{{key, value}}}}
			_, err = db.UpdateCoursePlan(filter, update, false)
			if err != nil {
				http.Error(w, err.Error(), 404)
				return
			}
		}
	}
}
package submit

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

func GetSubmissionByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["sid"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", objId}}
	sortby := bson.D{}
	problem, err := db.GetSingleSubmission(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "submission not found.", 404)
		return
	}
	util.ResponseJSON(w, problem)
}

func GetSubmissionsByUID(w http.ResponseWriter, r *http.Request) {
	user, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	uid := user.UserID
	filter := bson.D{{"creator", uid}}
	sortby := bson.D{{"createTime", -1}}

	pid_hex := r.URL.Query().Get("pid")
	if len(pid_hex) > 0 {
		// log.Println(pid_hex)
		pid, err := primitive.ObjectIDFromHex(pid_hex)
		if err != nil {
			// invalid id format
			log.Println(err)
		} else {
			filter = bson.D{{"$and", []bson.D{
				bson.D{{"creator", uid}},
				bson.D{{"problem", pid}},
			}}}
		}
	}

	submissions, err := db.GetMultipleSubmissions(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, err.Error(), 404)
		return
	}
	util.ResponseJSON(w, submissions)
}

func ReceiveJudgeResult(w http.ResponseWriter, r *http.Request) {
	sid, err := primitive.ObjectIDFromHex(mux.Vars(r)["sid"])
	if err != nil {
		// invalid sid format
		log.Println(err)
		return
	}
	caseNo, err := strconv.Atoi(mux.Vars(r)["caseNo"])
	if err != nil {
		// invalid case format
		log.Println(err)
		return
	}
	
	body, err := util.GetBody(r)
	result := schemas.Result{}
	_ = json.Unmarshal(body, &result)

	newJudgement := schemas.Judgement{
		TestcaseNo: caseNo, Token: result.Token, Status: result.Status.ID,
	}

	filter := bson.D{{"_id", sid}}
	update := bson.D{{"$push", bson.D{{"judgements", newJudgement}}}}
	_, err = db.UpdateSubmission(filter, update, false)

	update = bson.D{{"$inc", bson.D{{"judgementCompleted", 1}}}}
	_, err = db.UpdateSubmission(filter, update, false)

	update = bson.D{{"$bit", bson.D{{"status", bson.D{{"or", Status2Flag[result.Status.ID]}}}}}}
	_, err = db.UpdateSubmission(filter, update, false)
}

func GetMultipleSubmissions(w http.ResponseWriter, r *http.Request) {

}

func SubmitToProblem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["pid"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
		return
	}

	creator, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	problem, err := db.GetSingleProblem(bson.D{{"_id", objId}}, bson.D{})
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "problem not found.", 404)
		return
	}

	newSubmission := schemas.Submission{
		Creator: creator.UserID, Problem: problem.ProblemID, ProblemName: problem.Name, ProblemCategory: problem.Category,
		TestcaseCnt: problem.Testcase.TestcaseCnt, JudgementCompleted: 0, Judgements: []schemas.Judgement{}, CreateTime: time.Now(),
	}

	sid, err := db.CreateSubmission(newSubmission)

	body, err := util.GetBody(r)
	if err != nil {
		// http.Error()
	}
	userSubmission := schemas.UserSubmission{}
	err = json.Unmarshal(body, &userSubmission)
	if err != nil {
		// http.Error()
		return
	}
	// log.Println(userSubmission)

	err = SendJudgeRequests(problem, userSubmission, sid, r.URL.Query().Get("classroom"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 404)
		return
	}

	resp := struct {
		SubmissionID schemas.ID
	}{SubmissionID: sid}
	util.ResponseJSON(w, resp)
}

func ReceiveJudgeResult_Classroom(w http.ResponseWriter, r *http.Request) {
	crid, err := primitive.ObjectIDFromHex(mux.Vars(r)["crid"])
	if err != nil {
		// invalid sid format
		log.Println(err)
		return
	}
	sid, err := primitive.ObjectIDFromHex(mux.Vars(r)["sid"])
	if err != nil {
		// invalid sid format
		log.Println(err)
		return
	}
	caseNo, err := strconv.Atoi(mux.Vars(r)["caseNo"])
	if err != nil {
		// invalid case format
		log.Println(err)
		return
	}
	
	body, err := util.GetBody(r)
	result := schemas.Result{}
	_ = json.Unmarshal(body, &result)

	newJudgement := schemas.Judgement{
		TestcaseNo: caseNo, Token: result.Token, Status: result.Status.ID,
	}

	filter := bson.D{{"_id", sid}}
	update := bson.D{{"$push", bson.D{{"judgements", newJudgement}}}}
	_, err = db.UpdateSubmission(filter, update, false)

	update = bson.D{{"$inc", bson.D{{"judgementCompleted", 1}}}}
	_, err = db.UpdateSubmission(filter, update, false)

	update = bson.D{{"$bit", bson.D{{"status", bson.D{{"or", Status2Flag[result.Status.ID]}}}}}}
	_, err = db.UpdateSubmission(filter, update, false)

	sortby := bson.D{}
	submission, err := db.GetSingleSubmission(filter, sortby)
	if submission.JudgementCompleted < submission.TestcaseCnt {
		return
	}

	score := 0
	for _, judgement := range submission.Judgements {
		if judgement.Status == 3 {
			score += 1
		}
	}

	filter = bson.D{{"_id", crid}}
	update = bson.D{{"$addToSet", bson.D{{"scoreEntryList", schemas.ScoreEntry{UserID: submission.Creator, ProblemID: submission.Problem}}}}}
	_, err = db.UpdateClassroomRecord(filter, update, false)
	if err != nil {
		// failed to update score entry
		log.Println(err)
		return
	}

	filter = bson.D{{"$and", []bson.D{
		bson.D{{"_id", crid}},
		// bson.D{{"scoreEntryList", bson.D{{"$and", []bson.D{
		// 	bson.D{{"userID", submission.Creator}},
		// 	bson.D{{"problemID", submission.Problem}},
		// }}}}},
		bson.D{{"scoreEntryList.userID", submission.Creator}},
		bson.D{{"scoreEntryList.problemID", submission.Problem}},
	}}}
	update = bson.D{{"$max", bson.D{{"scoreEntryList.$.score", score}}}}
	_, err = db.UpdateClassroomRecord(filter, update, false)
	if err != nil {
		// failed to update score
		log.Println(err)
		return
	}
}

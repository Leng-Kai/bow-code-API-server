package submit

import (
	"encoding/json"
	"log"
	"net/http"
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

}

func ReceiveJudgeResult(w http.ResponseWriter, r *http.Request) {

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
		Creator: creator.UserID, Problem: problem.ProblemID, TestcaseCnt: problem.Testcase.TestcaseCnt,
		JudgementCompleted: 0, Judgements: []schemas.Judgement{}, CreateTime: time.Now(),
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
	log.Println(userSubmission)

	err = SendJudgeRequests(problem, userSubmission)
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

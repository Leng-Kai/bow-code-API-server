package problem

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

func GetAllProblems(w http.ResponseWriter, r *http.Request) {

}

func CreateNewProblem(w http.ResponseWriter, r *http.Request) {
	body, err := util.GetBody(r)
	if err != nil {
		// http.Error()
	}
	newProblem := schemas.Problem{}
	err = json.Unmarshal(body, &newProblem)
	if err != nil {
		// http.Error()
	}
	log.Println(newProblem)
	creator, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	newProblem.Creator = creator.UserID
	newProblem.CreateTime = time.Now()
	log.Println(newProblem)
	id, err := db.CreateProblem(newProblem)
	if err != nil {
		log.Println(err)
		// http.Error()
		return
	}
	_, err = db.UpdateUser(bson.D{{"_id", creator.UserID}}, bson.D{{"$push", bson.D{{"ownProblemList", id}}}}, true)
	if err != nil {
		log.Println(err)
		// http.Error()
		return
	}

	resp := struct {
		ProblemID schemas.ID
	}{ProblemID: id}
	util.ResponseJSON(w, resp)
}

func GetProblemByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["pid"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
		return
	}
	filter := bson.D{{"_id", objId}}
	sortby := bson.D{}
	problem, err := db.GetSingleProblem(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "problem not found.", 404)
		return
	}
	util.ResponseJSON(w, problem)
}

func UpdateProblemByID(w http.ResponseWriter, r *http.Request) {

}

func GetMultipleProblems(w http.ResponseWriter, r *http.Request) {
	param_tags := r.URL.Query().Get("tag")
	param_difficulties := r.URL.Query().Get("difficulty")
	filter_tags := bson.D{}
	filter_difficulties := bson.D{}

	if len(param_tags) > 0 {
		tags := strings.Split(param_tags, ",")
		filter_tags = bson.D{{
			"tags", bson.D{{
				"$elemMatch", bson.D{{
					"$in", tags,
				}},
			}},
		}}
	}
	if len(param_difficulties) > 0 {
		difficulties_string := strings.Split(param_difficulties, ",")
		difficulties := []int{}
		for _, difficulty_string := range difficulties_string {
			difficulty, _ := strconv.Atoi(difficulty_string)
			difficulties = append(difficulties, difficulty)
		}
		filter_difficulties = bson.D{{
			"difficulty", bson.D{{
				"$in", difficulties,
			}},
		}}
	}

	filter := bson.D{{"$and", []bson.D{
		filter_tags,
		filter_difficulties,
	}}}
	sortby := bson.D{}
	problems, err := db.GetMultipleProblems(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, err.Error(), 404)
		return
	}

	util.ResponseJSON(w, problems)
}

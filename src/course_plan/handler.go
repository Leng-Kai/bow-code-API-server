package course_plan

import (
	"encoding/json"
	"log"
	"net/http"
	// "strconv"
	// "strings"
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
		// http.Error()
		return
	}
	newCoursePlan := schemas.CoursePlan{}
	err = json.Unmarshal(body, &newCoursePlan)
	if err != nil {
		// http.Error()
		return
	}
	creator, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	newCoursePlan.Creator = creator.UserID
	newCoursePlan.CreateTime = time.Now()
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

}
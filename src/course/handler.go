package course

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Leng-Kai/bow-code-API-server/db"
	"github.com/Leng-Kai/bow-code-API-server/schemas"
	"github.com/Leng-Kai/bow-code-API-server/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{}
	sortby := bson.D{}
	allCourse, err := db.GetMultipleCourses(filter, sortby)
	if err != nil {
		//handle error
	}
	fmt.Print(allCourse)
	util.ResponseJSON(w, allCourse)
}

func CreateNew(w http.ResponseWriter, r *http.Request) {
	body, err := util.GetBody(r)
	if err != nil {
		// http.Error()
	}
	newCourse := schemas.Course{}
	err = json.Unmarshal(body, &newCourse)
	if err != nil {
		// http.Error()
	}
	id, err := db.CreateCourse(newCourse)
	if err != nil {
		//http.Error()
	} else {
		resp := struct {
			CourseID schemas.ID
		}{CourseID: id}
		util.ResponseJSON(w, resp)
	}
}

func GetCourseByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
	}
	filter := bson.D{{"_id", objId}}
	sortby := bson.D{}
	course, err := db.GetSingleCourse(filter, sortby)
	if err != nil {
		// db error
		log.Println(err)
		http.Error(w, "course not found.", 404)
		return
	}
	util.ResponseJSON(w, course)
}

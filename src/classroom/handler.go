package classroom

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

	resp := struct {
		ClassroomID schemas.ID
	}{ClassroomID: id}
	util.ResponseJSON(w, resp)
}

func ApplyForClassroom(w http.ResponseWriter, r *http.Request) {

}

func AcceptApplication(w http.ResponseWriter, r *http.Request) {

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
	util.ResponseJSON(w, classroom)
}

func UpdateClassroomByID(w http.ResponseWriter, r *http.Request) {

}
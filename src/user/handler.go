package user

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/Leng-Kai/bow-code-API-server/db"
	"github.com/Leng-Kai/bow-code-API-server/schemas"
	"github.com/Leng-Kai/bow-code-API-server/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := util.GetBody(r)
	if err != nil {
		// http.Error()
	}
	var msg map[string]string
	err = json.Unmarshal(body, &msg)
	if err != nil {
		// http.Error()
	}
	method := msg["method"]
	authenticator, ok := authHandler[method]
	if !ok {
		// unsupported register method
	}

	auth_payload := msg["authPayload"]
	auth_uid, userInfo, err := authenticator(auth_payload)
	if err != nil {
		// err may contain error message
	}

	global_uid := globalUid(method, auth_uid)

	// Check if user already exist
	filter := bson.D{{"_id", global_uid}}
	sortby := bson.D{}
	_, err = db.GetSingleUser(filter, sortby)
	if err != mongo.ErrNoDocuments {
		if err == nil {
			// user already exist
			log.Println("user already exist.")
			http.Error(w, "user already exist.", 404)
			return
		} else {
			// db error
		}
	}

	// If not, create a new user
	newUser := schemas.User{
		UserID: global_uid,
		RegisterType: method,
		UserInfo: userInfo,
		Super: false,
	}
	_, err = db.CreateUser(newUser)
	if err != nil {
		// db error
	} else {
		util.ResponseJSON(w, newUser)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := util.GetBody(r)
	if err != nil {
		// http.Error()
	}
	var msg map[string]string
	err = json.Unmarshal(body, &msg)
	if err != nil {
		// http.Error()
	}
	method := msg["method"]
	authenticator, ok := authHandler[method]
	if !ok {
		// unsupported register method
	}

	auth_payload := msg["authPayload"]
	auth_uid, userInfo, err := authenticator(auth_payload)
	if err != nil {
		// err may contain error message
	}

	global_uid := globalUid(method, auth_uid)

	// Check if user already exist
	filter := bson.D{{"_id", global_uid}}
	sortby := bson.D{}
	user, err := db.GetSingleUser(filter, sortby)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// user not exist
		} else {
			// db error
		}
	}

	// update the userInfo if needed
	if user.UserInfo != userInfo {
		filter := bson.D{{"_id", global_uid}}
		update := bson.D{{"$set", bson.D{{"userInfo", userInfo}}}}
		_, err = db.UpdateUser(filter, update, false)
		if err != nil {
			// db error
		}
		user.UserInfo = userInfo
	}

	util.ResponseJSON(w, user)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	filter := bson.D{{"_id", id}}
	sortby := bson.D{}
	user, err := db.GetSingleUser(filter, sortby)
	if err != nil {
		log.Println(err)
		http.Error(w, "user not found.", 404)
		return
	}
	util.ResponseJSON(w, user)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request){
	
}

package user

import (
	"net/http"

	"github.com/Leng-Kai/bow-code-API-server/db"
	"github.com/Leng-Kai/bow-code-API-server/schemas"
	"github.com/Leng-Kai/bow-code-API-server/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(w http.ResponseWriter, r *http.Request) {
	method := mux.Vars(r)["method"]
	authenticator, ok := authHandler[method]
	if !ok {
		// unsupported register method
	}

	auth_payload := mux.Vars(r)["authPayload"]
	auth_uid, userInfo, err := authenticator(auth_payload)
	if err != nil {
		// errr may contain error message
	}

	global_uid := globalUid(method, auth_uid)

	// Check if user already exist
	filter := bson.D{{"_id", global_uid}}
	sortby := bson.D{}
	user, err := db.GetSingleUser(filter, sortby)
	if err != nil {
		// db error
	}
	if user.UserID != "" {
		// user already exist
	}

	// If not, create a new user
	newUser := schemas.User{
		UserID: global_uid,
		RegisterType: method,
		UserInfo: userInfo,
		Super: false
	}
	_, err := db.CreateUser(newUser)
	if err != nil {
		// db error
	} else {
		resp := newUser
		util.ResponseJSON(w, resp)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func GetUserByID(w http.ResponseWriter, r *http.Request) {

}

func UpdateUserByID(w http.ResponseWriter, r *http.Request){
	
}

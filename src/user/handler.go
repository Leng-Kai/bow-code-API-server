package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Leng-Kai/bow-code-API-server/db"
	"github.com/Leng-Kai/bow-code-API-server/schemas"
	"github.com/Leng-Kai/bow-code-API-server/session"
	"github.com/Leng-Kai/bow-code-API-server/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		log.Println(err)
		http.Error(w, "unsupported register method.", 404)
		return
	}

	auth_payload := msg["authPayload"]
	auth_uid, userInfo, err := authenticator(auth_payload)
	if err != nil {
		// err may contain error message
		log.Println(err)
		http.Error(w, "failed to authenticate.", 404)
		return
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
		UserID:               global_uid,
		RegisterType:         method,
		UserInfo:             userInfo,
		SchoolInfo:           schemas.SchoolInfo{ Complete: false },
		Super:                false,
		JoinedCourseList:     []schemas.CourseID{},
		OwnCourseList:        []schemas.CourseID{},
		FavoriteCourseList:   []schemas.CourseID{},
		JoinedCoursePlanList: []schemas.CoursePlanID{},
		OwnCoursePlanList:    []schemas.CoursePlanID{},
		OwnProblemList:       []schemas.ProblemID{},
		FavoriteProblemList:  []schemas.ProblemID{},
		JoinedClassroomList:  []schemas.ClassroomID{},
		OwnClassroomList:     []schemas.ClassroomID{},
		AppliedClassroomList: []schemas.ClassroomID{},
	}
	_, err = db.CreateUser(newUser)

	if err != nil {
		// db error
	} else {
		session, err := session.Store.Get(r, "bow-session")
		if err != nil {
			log.Print(err)
		}
		session.Values["isLogin"] = true
		session.Values["uid"] = global_uid
		// Save it before we write to the response/return from the handler.
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
		log.Println(err)
		http.Error(w, "unsupported register method.", 404)
		return
	}

	auth_payload := msg["authPayload"]
	auth_uid, userInfo, err := authenticator(auth_payload)
	if err != nil {
		// err may contain error message
		log.Println(err)
		http.Error(w, "failed to authenticate.", 404)
		return
	}

	global_uid := globalUid(method, auth_uid)

	// Check if user already exist
	filter := bson.D{{"_id", global_uid}}
	sortby := bson.D{}
	user, err := db.GetSingleUser(filter, sortby)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// user not exist
			util.ResponseJSON(w, schemas.User{})
			return
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

	session, err := session.Store.Get(r, "bow-session")
	if err != nil {
		log.Print(err)
	}
	session.Values["isLogin"] = true
	session.Values["uid"] = global_uid
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.ResponseJSON(w, user)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := session.Store.Get(r, "bow-session")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 401)
		return
	}
	if islogin, ok := session.Values["isLogin"].(bool); ok && islogin {
		session.Values["isLogin"] = false
		delete(session.Values, "uid")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(200)
}

func AuthSession(w http.ResponseWriter, r *http.Request) {
	session, err := session.Store.Get(r, "bow-session")
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 401)
		return
	}

	if islogin, ok := session.Values["isLogin"].(bool); ok && islogin {
		filter := bson.D{{"_id", session.Values["uid"]}}
		sortby := bson.D{}
		user, err := db.GetSingleUser(filter, sortby)
		if err != nil {
			log.Println(err)
			http.Error(w, "user not found.", 404)
			return
		}
		util.ResponseJSON(w, user)
	} else {
		session.Values["isLogin"] = false
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		util.ResponseJSON(w, schemas.User{})
		return
	}

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

func UpdateSchoolInfo(w http.ResponseWriter, r *http.Request) {
	body, err := util.GetBody(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	newSchoolInfo := schemas.SchoolInfo{}
	err = json.Unmarshal(body, &newSchoolInfo)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	newSchoolInfo.Complete = true
	user_obj, err := GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	_, err = db.UpdateUser(bson.D{{"_id", uid}}, bson.D{{"$set", bson.D{{"schoolInfo", newSchoolInfo}}}}, false)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {

}

package bulletin

import (
	// "encoding/json"
	// "log"
	"net/http"
	// "strconv"
	// "strings"
	// "time"

	// . "github.com/Leng-Kai/bow-code-API-server/course_plan"
	"github.com/Leng-Kai/bow-code-API-server/db"
	// "github.com/Leng-Kai/bow-code-API-server/schemas"
	"github.com/Leng-Kai/bow-code-API-server/user"
	// "github.com/Leng-Kai/bow-code-API-server/util"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LikeBulletin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["bid"]
	bid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	filter := bson.D{{"_id", bid}}
	update := bson.D{{"$addToSet", bson.D{{"reactions", uid}}}}
	_, err = db.UpdateBulletin(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
}

func UnlikeBulletin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["bid"]
	bid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	filter := bson.D{{"_id", bid}}
	update := bson.D{{"$pull", bson.D{{"reactions", uid}}}}
	_, err = db.UpdateBulletin(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
}

func DeleteBulletin(w http.ResponseWriter, r *http.Request) {
	crid, err := primitive.ObjectIDFromHex(mux.Vars(r)["crid"])
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	bid, err := primitive.ObjectIDFromHex(mux.Vars(r)["bid"])
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	filter := bson.D{{"_id", bid}}
	sortby := bson.D{}
	bulletin, err := db.GetSingleBulletin(filter, sortby)
	if err != nil {
		http.Error(w, "bulletin not found.", 404)
		return
	}

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID
	if uid != bulletin.Creator {
		http.Error(w, "permission denied. not bulletin creator.", 404)
		return
	}

	_, err = db.DeleteBulletin(filter, bson.D{})
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	filter = bson.D{{"_id", crid}}
	update := bson.D{{"$pull", bson.D{{"bulletinList", bid}}}}
	_, err = db.UpdateClassroom(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
}
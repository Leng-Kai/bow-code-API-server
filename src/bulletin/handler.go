package bulletin

import (
	"encoding/json"
	// "log"
	"net/http"
	"strconv"
	// "strings"
	"time"

	// . "github.com/Leng-Kai/bow-code-API-server/course_plan"
	"github.com/Leng-Kai/bow-code-API-server/db"
	"github.com/Leng-Kai/bow-code-API-server/schemas"
	"github.com/Leng-Kai/bow-code-API-server/user"
	"github.com/Leng-Kai/bow-code-API-server/util"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewBulletin(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["crid"]
	crid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	filter := bson.D{{"_id", crid}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		http.Error(w, "classroom not found.", 404)
		return
	}

	creator, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	if creator.UserID != classroom.Creator && !util.Contain_str(classroom.Students, creator.UserID) {
		http.Error(w, "not in classroom.", 404)
		return
	}

	body, err := util.GetBody(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	newBulletin := schemas.Bulletin{}
	err = json.Unmarshal(body, &newBulletin)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	newBulletin.Classroom = crid
	newBulletin.Creator = creator.UserID
	newBulletin.Reactions = []schemas.UserID{}
	newBulletin.Replies = []schemas.Reply{}
	newBulletin.IndexCount = 0
	newBulletin.CreateTime = time.Now()

	bid, err := db.CreateBulletin(newBulletin)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	_, err = db.UpdateClassroom(bson.D{{"_id", crid}}, bson.D{{"$push", bson.D{{"bulletinList", bid}}}}, false)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	resp := struct {
		BulletinID schemas.ID
	}{BulletinID: bid}
	util.ResponseJSON(w, resp)
}

func LikeReply(w http.ResponseWriter, r *http.Request) {
	bid, err := primitive.ObjectIDFromHex(mux.Vars(r)["bid"])
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	index, err := strconv.Atoi(mux.Vars(r)["index"])
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

	filter := bson.D{{"$and", []bson.D{
		bson.D{{"_id", bid}},
		bson.D{{"replies.index", index}},
	}}}
	update := bson.D{{"$addToSet", bson.D{{"replies.$.reactions", uid}}}}
	_, err = db.UpdateBulletin(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
}

func ReplyToBulletin(w http.ResponseWriter, r *http.Request) {
	bid, err := primitive.ObjectIDFromHex(mux.Vars(r)["bid"])
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	filter := bson.D{{"_id", bid}}
	update := bson.D{{"$inc", bson.D{{"indexCount", 1}}}}
	oldBulletin, err := db.UpdateBulletin(filter, update, false)
	if err != nil {
		http.Error(w, "bulletin not found.", 404)
		return
	}
	index := oldBulletin.IndexCount

	user_obj, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	uid := user_obj.UserID

	body, err := util.GetBody(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	newReply := schemas.Reply{}
	err = json.Unmarshal(body, &newReply)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	newReply.Creator = uid
	newReply.Reactions = []schemas.UserID{}
	newReply.Index = index
	newReply.CreateTime = time.Now()

	update = bson.D{{"$addToSet", bson.D{{"replies", newReply}}}}
	_, err = db.UpdateBulletin(filter, update, false)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
}

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
	crid := bulletin.Classroom

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

func EditBulletin(w http.ResponseWriter, r *http.Request) {
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

	body, err := util.GetBody(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	var updatedBulletin map[string]interface{}
	err = json.Unmarshal(body, &updatedBulletin)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	legal_key := []string{"title", "content"}
	for key, value := range updatedBulletin {
		if util.Contain_str(legal_key, key) {
			update := bson.D{{"$set", bson.D{{key, value}}}}
			_, err = db.UpdateBulletin(filter, update, false)
			if err != nil {
				http.Error(w, err.Error(), 401)
				return
			}
		}
	}
}
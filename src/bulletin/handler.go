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
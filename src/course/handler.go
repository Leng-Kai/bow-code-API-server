package course

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
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
	docsPath := os.Getenv("DOCS_PATH")
	os.Mkdir(path.Join(docsPath, "course"), 0755)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	filter := bson.D{}
	sortby := bson.D{}

	for k, v := range params {
		filter = append(filter, bson.E{k, bson.D{{"$in", v}}})
	}
	allCourse, err := db.GetMultipleCourses(filter, sortby)
	if err != nil {
		//handle error
	}
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

	creator, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}
	newCourse.Creator = creator.UserID
	newCourse.CreateTime = time.Now()
	newCourse.Views = 0
	id, err := db.CreateCourse(newCourse)
	if err != nil {
		//http.Error()
	} else {
		_, err := db.UpdateUser(bson.D{{"_id", creator.UserID}}, bson.D{{"$push", bson.D{{"ownCourseList", id}}}}, true)
		if err != nil {
			log.Println(err)
		}
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

func GetMultipleCourses(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	coursesId := []schemas.CourseID{}

	for _, id := range params["courses"] {
		objId, err := primitive.ObjectIDFromHex(id)

		if err == nil {
			coursesId = append(coursesId, objId)
		}
	}
	filter := bson.D{{"_id", bson.D{{"$in", coursesId}}}}
	sortby := bson.D{}

	courses, err := db.GetMultipleCourses(filter, sortby)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	util.ResponseJSON(w, courses)
}

func UpdateCourseByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
	}
	creator, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
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

	if course.Creator != creator.UserID {
		http.Error(w, "permission denied. not course creator.", 401)
		return
	}

	body, err := util.GetBody(r)
	if err != nil {
		// http.Error()
	}
	newCourse := schemas.Course{}
	err = json.Unmarshal(body, &newCourse)
	if err != nil {
		// http.Error()
	}
	newCourse.BlockList = course.BlockList // NO BLOCKLIST CHANGE HERE!!
	err = db.ReplaceCourse(filter, newCourse)
	if err != nil {
		http.Error(w, "update failed", 400)
		return
	}
	w.WriteHeader(200)
}

func RemoveCourseByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
	}
	creator, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
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

	if course.Creator != creator.UserID {
		http.Error(w, "permission denied. not course creator.", 401)
		return
	}

	projection := bson.D{{"_id", 1}}
	_, err = db.DeleteCourse(filter, projection)
	if err != nil {
		log.Println(err)
		http.Error(w, "delete failed", 400)
		return
	}
	util.ResponseJSON(w, course)
}

func LoveCourseByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
	}
	user, err := user.GetSessionUser(r)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	_, err = db.UpdateUser(bson.D{{"_id", user.UserID}}, bson.D{{"$addToSet", bson.D{{"favoriteCourseList", objId}}}}, true)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func CreateBlock(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
	}
	docsPath := os.Getenv("DOCS_PATH")
	blockContent, err := util.GetBody(r)
	if err != nil {
		// http.Error()
	}

	files, err := ioutil.ReadDir(path.Join(docsPath, "course", id, "block"))
	newBlockPath := path.Join(docsPath, "course", id, "block")

	/** Make directory if dir not exist **/
	if _, err := os.Stat(newBlockPath); os.IsNotExist(err) {
		err = os.MkdirAll(newBlockPath, 0755)
		if err != nil {
			log.Println(err)
		}
	}

	/** Retrieve the new block ID **/
	/** Some error will occur when block ID exceed 99999 **/
	newBlockID := 10000
	if len(files) > 0 {
		lastID, err := strconv.Atoi(files[len(files)-1].Name())
		if err != nil {
			// invalid file name
		}
		newBlockID = lastID + 1
	}
	newBlockPath = path.Join(newBlockPath, strconv.Itoa(newBlockID))
	err = ioutil.WriteFile(newBlockPath, blockContent, 0644)
	if err != nil {
		// Error to write to file
		log.Println(err)
	}

	sc := bufio.NewScanner(strings.NewReader(string(blockContent)))
	sc.Scan()
	title := sc.Text()

	filter := bson.D{{"_id", objId}}
	blockEntry := bson.D{{"title", title}, {"ID", strconv.Itoa(newBlockID)}}
	update := bson.D{{"$push", bson.D{{"blockList", blockEntry}}}}
	_, err = db.UpdateCourse(filter, update, false)
	if err != nil {
		// update failed
	}
	w.WriteHeader(200)
}

func GetBlock(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	docsPath := os.Getenv("DOCS_PATH")
	blockId := mux.Vars(r)["bid"]
	blockContent, err := ioutil.ReadFile(path.Join(docsPath, "course", id, "block", blockId))
	if err != nil {
		// failed to read file
	}
	util.ResponseHTML(w, string(blockContent))
}

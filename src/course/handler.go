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
	filter := bson.D{}
	sortby := bson.D{}
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
		log.Println(err)
		// http.Error()
	} else {
		docs_path := os.Getenv("DOCS_PATH")
		newBlockPath := path.Join(docs_path, "course", id.Hex(), "block")
		err = os.MkdirAll(newBlockPath, 0777)
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

func CreateBlock(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// invalid id format
		log.Println(err)
	}
	docsPath := os.Getenv("DOCS_PATH")
	// blockContent, err := util.GetBody(r)
	// if err != nil {
	// 	// http.Error()
	// }

	files, err := ioutil.ReadDir(path.Join(docsPath, "course", id, "block"))

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

	newBlockPath := path.Join(docsPath, "course", id, "block", strconv.Itoa(newBlockID))

	err = os.MkdirAll(newBlockPath, 0777)
	if err != nil {
		log.Println(err)
	}

	_, err = os.Create(path.Join(newBlockPath, "index.html"))
	if err != nil {
		log.Println(err)
	}

	// err = ioutil.WriteFile(newBlockPath, blockContent, 0644)
	// if err != nil {
	// 	// Error to write to file
	// 	log.Println(err)
	// }

	// sc := bufio.NewScanner(strings.NewReader(string(blockContent)))
	// sc.Scan()
	// title := sc.Text()
	title := ""

	filter := bson.D{{"_id", objId}}
	blockEntry := bson.D{{"title", title}, {"ID", strconv.Itoa(newBlockID)}}
	update := bson.D{{"$push", bson.D{{"blockList", blockEntry}}}}
	_, err = db.UpdateCourse(filter, update, false)
	if err != nil {
		// update failed
	}

	// w.WriteHeader(200)
	util.ResponseJSON(w, strconv.Itoa(newBlockID))
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

func UpdateBlock(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bid := mux.Vars(r)["bid"]
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

	docsPath := os.Getenv("DOCS_PATH")
	blockContent, err := util.GetBody(r)
	if err != nil {
		// http.Error()
	}

	blockPath := path.Join(docsPath, "course", id, "block", bid, "index.html")

	err = ioutil.WriteFile(blockPath, blockContent, 0777)
	if err != nil {
		// Error to write to file
		log.Println(err)
	}

	sc := bufio.NewScanner(strings.NewReader(string(blockContent)))
	sc.Scan()
	title := sc.Text()

	filter = bson.D{{"_id", objId}}
	blockEntry := bson.D{{"title", title}, {"ID", bid}}
	update := bson.D{{"$push", bson.D{{"blockList", blockEntry}}}}
	_, err = db.UpdateCourse(filter, update, false)
	if err != nil {
		// update failed
	}
}

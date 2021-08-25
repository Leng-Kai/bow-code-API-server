package classroom

import (
	"github.com/Leng-Kai/bow-code-API-server/course_plan"
	"github.com/Leng-Kai/bow-code-API-server/db"
	"github.com/Leng-Kai/bow-code-API-server/schemas"

	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	
}

func AddRecordsForNewStudent(crid schemas.ClassroomID, uid schemas.UserID) error {
	filter := bson.D{{"_id", crid}}
	sortby := bson.D{}
	classroom, err := db.GetSingleClassroom(filter, sortby)
	if err != nil {
		return err
	}

	cpid := classroom.CoursePlan
	filter = bson.D{{"_id", cpid}}
	coursePlan, err := db.GetSingleCoursePlan(filter, sortby)
	if err != nil {
		return err
	}

	componentList := coursePlan.ComponentList
	scoreEntries := []schemas.ScoreEntry{}

	for _, component := range componentList {
		if component.Type == course_plan.PROBLEM {
			scoreEntry := schemas.ScoreEntry{ UserID: uid, ProblemID: component.ID, Score: -1 }
			scoreEntries = append(scoreEntries, scoreEntry)
		}
	}

	update := bson.D{{"$push", bson.D{{"scoreEntryList", bson.D{{"$each", scoreEntries}}}}}}
	_, err = db.UpdateClassroomRecord(filter, update, false)
	return err
}

// func AddRecordsForNewStudent(crid schemas.ClassroomID, uid schemas.UserID) error {
// 	filter := bson.D{{"_id", crid}}
// 	sortby := bson.D{}
// 	classroom, err := db.GetSingleClassroom(filter, sortby)
// 	if err != nil {
// 		return err
// 	}

// 	cpid := classroom.CoursePlan
// 	filter = bson.D{{"_id", cpid}}
// 	coursePlan, err := db.GetSingleCoursePlan(filter, sortby)
// 	if err != nil {
// 		return err
// 	}

// 	componentList := coursePlan.ComponentList
// 	scoreEntries := []schemas.ScoreEntry{}

// 	for _, component := range componentList {
// 		if component.Type == course_plan.PROBLEM {
// 			scoreEntry := schemas.ScoreEntry{ UserID: uid, ProblemID: component.ID, Score: -1 }
// 			scoreEntries = append(scoreEntries, scoreEntry)
// 		}
// 	}

// 	update := bson.D{{"$push", bson.D{{"scoreEntryList", bson.D{{"$each", scoreEntries}}}}}}
// 	_, err = db.UpdateClassroomRecord(filter, update, false)
// 	return err
// }
package schemas

import (
	"time"
)

type CoursePlanID = ID

type CoursePlan struct {
	CoursePlanID CoursePlanID `json:"id" bson:"_id,omitempty"`
	Name         string       `json:"name" bson:"name"`
	CourseList   []CourseID   `json:"courseList" bson:"courseList"`
	Creator      UserID       `json:"creator" bson:"creator"`
	Visibility   int          `json:"visibility" bson:"visibility"`
	CreateTime   time.Time    `json:"createTime" bson:"createTime"`
}

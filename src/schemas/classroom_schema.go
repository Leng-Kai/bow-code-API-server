package schemas

import (
	"time"
)

type ClassroomID = ID

type Classroom struct {
	ClassroomID ClassroomID  `json:"id" bson:"_id,omitempty"`
	Name        string       `json:"name" bson:"name"`
	CoursePlan  CoursePlanID `json:"coursePlan" bson:"coursePlan"`
	Creator     UserID       `json:"creator" bson:"creator"`
	Students    []UserID     `json:"students" bson:"students"`
	Visibility  int          `json:"visibility" bson:"visibility"`
	CreateTime  time.Time    `json:"createTime" bson:"createTime"`
}

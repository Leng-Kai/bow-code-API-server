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
	Review      bool         `json:"review" bson:"review"`
	Apply       bool         `json:"apply" bson:"apply"`
	Applicants  []UserID     `json:"applicants" bson:"applicants"`
	Invitees    []UserID     `json:"invitees" bson:"invitees"`
	Visibility  int          `json:"visibility" bson:"visibility"`
	CreateTime  time.Time    `json:"createTime" bson:"createTime"`
}

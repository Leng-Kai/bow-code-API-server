package schemas

import (
	"time"
)

type ClassroomID = ID

type ClassroomComponent struct {
	Component CoursePlanComponent `json:"component" bson:"component"`
	Begin     int                 `json:"begin" bson:"begin"`
	End       int                 `json:"end" bson:"end"`
	Private   bool                `json:"private" bson:"private"`
}

type Classroom struct {
	ClassroomID  ClassroomID          `json:"id" bson:"_id,omitempty"`
	Name         string               `json:"name" bson:"name"`
	CoursePlan   CoursePlanID         `json:"coursePlan" bson:"coursePlan"`
	Creator      UserID               `json:"creator" bson:"creator"`
	Students     []UserID             `json:"students" bson:"students"`
	Review       bool                 `json:"review" bson:"review"`
	Apply        bool                 `json:"apply" bson:"apply"`
	Applicants   []UserID             `json:"applicants" bson:"applicants"`
	Invitees     []UserID             `json:"invitees" bson:"invitees"`
	BulletinList []BulletinID         `json:"bulletinList" bson:"bulletinList"`
	CourseList   []ClassroomComponent `json:"courseList" bson:"courseList"`
	HomeworkList []ClassroomComponent `json:"homeworkList" bson:"homeworkList"`
	ExamList     []ClassroomComponent `json:"examList" bson:"examList"`
	Visibility   int                  `json:"visibility" bson:"visibility"`
	CreateTime   time.Time            `json:"createTime" bson:"createTime"`
}

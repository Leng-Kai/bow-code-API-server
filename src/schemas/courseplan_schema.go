package schemas

import (
	"time"
)

type CoursePlanID = ID

type CoursePlanComponent struct {
	Type int `json:"type" bson:"type"`
	ID   ID  `json:"id" bson:"id"`
}

type CoursePlan struct {
	CoursePlanID  CoursePlanID          `json:"id" bson:"_id,omitempty"`
	Name          string                `json:"name" bson:"name"`
	ComponentList []CoursePlanComponent `json:"componentList" bson:"componentList"`
	Creator       UserID                `json:"creator" bson:"creator"`
	Visibility    int                   `json:"visibility" bson:"visibility"`
	CreateTime    time.Time             `json:"createTime" bson:"createTime"`
}

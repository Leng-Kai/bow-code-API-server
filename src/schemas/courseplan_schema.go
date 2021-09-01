package schemas

import (
	"time"
)

type CoursePlanID = ID

type Set struct {
	Name       string `json:"name" bson:"name"`
	ID         ID     `json:"id" bson:"id"`
	TotalScore int    `json:"totalScore" bson:"totalScore"`
}

type CoursePlanComponent struct {
	Name    string `json:"name" bson:"name"`
	Type    int    `json:"type" bson:"type"`
	SetList []Set  `json:"setList" bson:"setList"`
}

type CoursePlan struct {
	CoursePlanID  CoursePlanID          `json:"id" bson:"_id,omitempty"`
	Name          string                `json:"name" bson:"name"`
	ComponentList []CoursePlanComponent `json:"componentList" bson:"componentList"`
	Creator       UserID                `json:"creator" bson:"creator"`
	Visibility    int                   `json:"visibility" bson:"visibility"`
	CreateTime    time.Time             `json:"createTime" bson:"createTime"`
}

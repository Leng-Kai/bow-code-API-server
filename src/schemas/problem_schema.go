package schemas

import (
	"time"
)

type ProblemID = ID
type Testcase struct {
	TestcaseCnt    int      `json:"testcaseCnt", bson:"testcaseCnt`
	Input          []string `json:"input" bson:"input"`
	ExpectedOutput []string `json:"expectedOutput" bson:"expectedOutput"`
}

type Problem struct {
	ProblemID      ProblemID `json:"id" bson:"_id,omitempty"`
	Name           string    `json:"name" bson:"name"`
	Creator        UserID    `json:"creator" bson:"creator"`
	Description    string    `json:"description" bson:"description"`
	DefaultContent string    `json:"defaultContent" bson:"defaultContent"`
	Testcase       Testcase  `json:"testcase" bson:"testcase"`
	Tags           []string  `json:"tags" bson:"tags"`
	Difficulty     int       `json:"difficulty" bson:"difficulty"`
	Category       string    `json:"category" bson:"category"`
	// IsPublic       bool      `json:"isPublic" bson:"isPublic"`
	Visibility int       `json:"visibility" bson:"visibility"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
}

type TagCount struct {
	Tag   string `json:"tag" bson:"tag"`
	Count int    `json:"count" bson:"count"`
}

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
	ProblemID  ID        `json:"id" bson:"_id,omitempty"`
	Name       string    `json:"name" bson:"name"`
	Creator    ID        `json:"creator" bson:"creator"`
	Testcase   Testcase  `json:"testcase" bson:"testcase"`
	Tags       []string  `json:"tags" bson:"tags"`
	Difficulty int       `json:"difficulty" bson:"difficulty"`
	Category   string    `json:"category" bson:"category"`
	IsPublic   bool      `json:"isPublic" bson:"isPublic"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
}

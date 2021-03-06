package schemas

import (
	"time"
)

type ProblemID = ID

type Testcase struct {
	TestcaseCnt    int      `json:"testcaseCnt", bson:"testcaseCnt`
	Input          []string `json:"input" bson:"input"`
	ExpectedOutput []string `json:"expectedOutput" bson:"expectedOutput"`
	Score          []int    `json:"score" bson:"score"`
	ShowDetail     []bool   `json:"showDetail" bson:"showDetail"`
}

type Contents struct {
	Language int    `json:"language" bson:"language"`
	Content  string `json:"content" bson:"content"`
}

type Problem struct {
	ProblemID      ProblemID  `json:"id" bson:"_id,omitempty"`
	Name           string     `json:"name" bson:"name"`
	Creator        UserID     `json:"creator" bson:"creator"`
	Description    string     `json:"description" bson:"description"`
	DefaultContent []Contents `json:"defaultContent" bson:"defaultContent"`
	Testcase       Testcase   `json:"testcase" bson:"testcase"`
	TotalScore     int        `json:"totalScore" bson:"totalScore"`
	Tags           []string   `json:"tags" bson:"tags"`
	Difficulty     int        `json:"difficulty" bson:"difficulty"`
	Category       string     `json:"category" bson:"category"`
	Visibility     int        `json:"visibility" bson:"visibility"`
	CreateTime     time.Time  `json:"createTime" bson:"createTime"`
}

type TagCount struct {
	Tag   string `json:"tag" bson:"tag"`
	Count int    `json:"count" bson:"count"`
}

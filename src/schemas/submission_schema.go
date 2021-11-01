package schemas

import (
	"time"
)

type SubmissionID = ID
type SubmissionToken = string

type UserSubmission struct {
	SourceCode string `json:"sourceCode" bson:"sourceCode"`
	LanguageID int    `json:"languageID" bson:"languageID"`
}

type Judgement struct {
	TestcaseNo      int             `json:"testcaseNo" bson:"testcaseNo"`
	Input           string          `json:"input" bson:"input"`
	Expected_output string          `json:"expected_output" bson:"expected_output"`
	Output          string          `json:"output" bson:"output"`
	Token           SubmissionToken `json:"token" bson:"token"`
	Status          int             `json:"status" bson:"status"`
}

type Submission struct {
	SubmissionID       ID          `json:"id" bson:"_id,omitempty"`
	Creator            UserID      `json:"creator" bson:"creator"`
	Problem            ID          `json:"problem" bson:"problem"`
	ProblemName        string      `json:"problemName" bson:"problemName"`
	ProblemCategory    string      `json:"problemCategory" bson:"problemCategory"`
	TestcaseCnt        int         `json:"testcaseCnt" bson:"testcaseCnt"`
	JudgementCompleted int         `json:"judgementCompleted" bson:"judgementCompleted"`
	Judgements         []Judgement `json:"judgements" bson:"judgements"`
	Status             uint        `json:"status" bson:"status"`
	CreateTime         time.Time   `json:"createTime" bson:"createTime"`
}

type Status struct {
	Description string `json:"description"`
	ID          int    `json:"id"`
}

type Result struct {
	Compile_output  string  `json:"compile_output"`
	Memory          int     `json:"memory"`
	Message         string  `json:"message"`
	Status          Status  `json:"status"`
	Stdin           string  `json:"stdin"`
	Stderr          string  `json:"stderr"`
	Stdout          string  `json:"stdout"`
	Expected_output string  `json:"expected_output"`
	Time            float32 `json:"time"`
	Token           string  `json:"token"`
}

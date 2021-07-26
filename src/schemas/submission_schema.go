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
	TestcaseNo int             `json:"testcaseNo" bson:"testcaseNo"`
	Token      SubmissionToken `json:"token" bson:"token"`
	Status     int             `json:"status" bson:"status"`
}

type Submission struct {
	SubmissionID       ID          `json:"id" bson:"_id,omitempty"`
	Creator            UserID      `json:"creator" bson:"creator"`
	Problem            ID          `json:"problem" bson:"problem"`
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
	Compile_output string  `json:"compile_output"`
	Memory         int     `json:"memory"`
	Message        string  `json:"message"`
	Status         Status  `json:"status"`
	Stderr         string  `json:"stderr"`
	Stdout         string  `json:"stdout"`
	Time           float32 `json:"time"`
	Token          string  `json:"token"`
}

package schemas

import (
	"time"
)

type SubmissionID = ID
type SubmissionToken = string

type Judgement struct {
	TestcaseNo int             `json:"testcaseNo" bson:"testcaseNo"`
	Token      SubmissionToken `json:"token" bson:"token"`
}

type Submission struct {
	SubmissionID       ID          `json:"id" bson:"_id,omitempty"`
	Creator            ID          `json:"creator" bson:"creator"`
	Problem            ID          `json:"problem" bson:"problem"`
	TestcaseCnt        int         `json:"testcaseCnt" bson:"testcaseCnt"`
	JudgementCompleted int         `json:"judgementCompleted" bson:"judgementCompleted"`
	Judgements         []Judgement `json:"judgements" bson:"judgements"`
	CreateTime         time.Time   `json:"createTime" bson:"createTime"`
}

package schemas

import (
	"time"
)

type SubmissionID = ID
type SubmissionToken = string

type UserSubmission struct {
	SourceCode string `json:"sourceCode" bson:"sourceCode"`
	LanguageID int	  `json:"languageID" bson:"languageID"`
}

type Judgement struct {
	TestcaseNo int             `json:"testcaseNo" bson:"testcaseNo"`
	Token      SubmissionToken `json:"token" bson:"token"`
	Status	   int 			   `json:"status" bson:"status"`
}

type Submission struct {
	SubmissionID       ID          `json:"id" bson:"_id,omitempty"`
	Creator            UserID      `json:"creator" bson:"creator"`
	Problem            ID          `json:"problem" bson:"problem"`
	TestcaseCnt        int         `json:"testcaseCnt" bson:"testcaseCnt"`
	JudgementCompleted int         `json:"judgementCompleted" bson:"judgementCompleted"`
	Judgements         []Judgement `json:"judgements" bson:"judgements"`
	Status	   		   int 		   `json:"status" bson:"status"`
	CreateTime         time.Time   `json:"createTime" bson:"createTime"`
}

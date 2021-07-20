package schemas

import (
	"time"
)

type SubmissionID = ID
type SubmissionToken = string

type Submission struct {
	SubmissionID ID              `json:"id" bson:"_id,omitempty"`
	Token        SubmissionToken `json:"token" bson:"token"`
	Creator      ID              `json:"creator" bson:"creator"`
	Problem      ID              `json:"problem" bson:"problem"`
	TestcaseNo   int             `json:"testcaseNo" bson:"testcaseNo"`
	CreateTime   time.Time       `json:"createTime" bson:"createTime"`
}

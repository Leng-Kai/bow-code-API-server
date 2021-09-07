package schemas

import (
	"time"
)

type BulletinID = ID

type Reply struct {
	Index      int       `json:"index" bson:"index"`
	Creator    UserID    `json:"creator" bson:"creator"`
	Content    string    `json:"content" bson:"content"`
	Reactions  []UserID  `json:"reactions" bson:"reactions"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
}

type Bulletin struct {
	BulletinID BulletinID `json:"id" bson:"_id,omitempty"`
	Classroom  ClassroomID `json:"classroom" bson:"classroom"`
	Creator    UserID     `json:"creator" bson:"creator"`
	Title      string     `json:"title" bson:"title"`
	Content    string     `json:"content" bson:"content"`
	Reactions  []UserID   `json:"reactions" bson:"reactions"`
	Replies    []Reply    `json:"replies" bson:"replies"`
	IndexCount int        `json:"indexCount" bson:"indexCount"`
	CreateTime time.Time  `json:"createTime" bson:"createTime"`
}

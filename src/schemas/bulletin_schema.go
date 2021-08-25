package schemas

import (
	"time"
)

type BulletinID = ID
type BulletinBoardID = ID

type Reply struct {
	Creator    UserID    `json:"creator" bson:"creator"`
	Reactions  []UserID  `json:"reactions" bson:"reactions"`
	Content    string    `json:"content" bson:"content"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
}

type Bulletin struct {
	BulletinID BulletinID `json:"id" bson:"_id,omitempty"`
	Creator    UserID     `json:"creator" bson:"creator"`
	Title      string     `json:"title" bson:"title"`
	Content    string     `json:"content" bson:"content"`
	Reactions  []UserID   `json:"reactions" bson:"reactions"`
	Replies    []Reply    `json:"replies" bson:"replies"`
	CreateTime time.Time  `json:"createTime" bson:"createTime"`
}

type BulletinBoard struct {
	BulletinBoardID BulletinBoardID `json:"id" bson:"_id,omitempty"`
	BulletinList    []BulletinID    `json:"bulletinList" bson:"bulletinList"`
}

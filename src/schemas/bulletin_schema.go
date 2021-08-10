package schemas

import (
	"time"
)

type BulletinID = ID
type BulletinBoardID = ID

type Reply struct {
	Creator     UserID    `json:"creator" bson:"creator"`
	CreatorInfo UserInfo  `json:"creatorInfo" bson:"creatorInfo"`
	Reactions   []UserID  `json:"reactions" bson:"reactions"`
	Content     string    `json:"content" bson:"content"`
	CreateTime  time.Time `json:"createTime" bson:"createTime"`
}

type Bulletin struct {
	BulletinID  BulletinID `json:"id" bson:"_id,omitempty"`
	Creator     UserID     `json:"creator" bson:"creator"`
	CreatorInfo UserInfo   `json:"creatorInfo" bson:"creatorInfo"`
	Title       string     `json:"title" bson:"title"`
	Content     string     `json:"content" bson:"content"`
	Reactions   []UserID   `json:"reactions" bson:"reactions"`
	Replies     []Reply    `json:"replies" bson:"replies"`
	CreateTime  time.Time  `json:"createTime" bson:"createTime"`
}

type BulletinBoard struct {
	BulletinBoardID BulletinBoardID `json:"id" bson:"_id,omitempty"`
	BulletinList    []Bulletin      `json:"bulletinList" bson:"bulletinList"`
}

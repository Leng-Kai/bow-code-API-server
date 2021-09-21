package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ID = primitive.ObjectID
type CourseBlockID = ID
type CourseID = ID

type Filter = bson.D
type Sortby = bson.D
type Projection = bson.D
type Update = bson.D

type CourseBlock struct {
	CourseBlockID ID `json:"_id,omitempty"`
}

type Course struct {
	CourseID  ID     `json:"id" bson:"_id,omitempty"`
	Name      string `json:"name" bson:"name"`
	Abstract  string `json:"abstract" bson:"abstract"`
	Image     string `json:"image" bson:"image"`
	BlockList []struct {
		Type  string `json:"type"`
		Title string `json:"title"`
		ID    string `json:"id"`
	} `json:"blockList" bson:"blockList"`
	Creator    UserID    `json:"creator" bson:"creator"`
	Tags       []string  `json:"tags" bson:"tags"`
	Difficulty int       `json:"difficulty" bson:"difficulty"`
	Category   string    `json:"category" bson:"category"`
	IsPublic   bool      `json:"isPublic" bson:"isPublic"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
	Views      int       `json:"views" bson:"views"`
}

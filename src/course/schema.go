package course

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ID 			= primitive.ObjectID
type CourseBlockID 	= ID
type CourseID 		= ID
type CoursePlanID 	= ID

type Filter			= bson.D
type Sortby			= bson.D
type Projection 	= bson.D
type Update			= bson.D

type CourseBlock struct {
	CourseBlockID			`json:"id" 			bson:"_id, omitempty"` 
}

type Course struct {
	CourseID				`json: "id" 		bson: "_id, omitempty"` 
	BlockList	[]ID		`json: "blockList" 	bson: "blockList"`
	Creator   	ID			`json: "creator"	bson: "creator"`
}

type CoursePlan struct {
	CoursePlanID			`json:"id" 			bson:"_id, omitempty"` 
}

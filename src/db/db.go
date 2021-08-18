package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const db_name = "bow-code-DB"

var db *mongo.Client
var courses *mongo.Collection
var courseplans *mongo.Collection
var classrooms *mongo.Collection
var classroomrecords *mongo.Collection
var users *mongo.Collection
var problems *mongo.Collection
var submissions *mongo.Collection
// var Session *mongo.Collection

func InitDB(client *mongo.Client) {
	db = client
	courses = db.Database(db_name).Collection("courses")
	courseplans = db.Database(db_name).Collection("courseplans")
	classrooms = db.Database(db_name).Collection("classrooms")
	classroomrecords = db.Database(db_name).Collection("classroomrecords")
	users = db.Database(db_name).Collection("users")
	problems = db.Database(db_name).Collection("problems")
	submissions = db.Database(db_name).Collection("submissions")
	// Session = db.Database("SessionDB").Collection("sessions")
}

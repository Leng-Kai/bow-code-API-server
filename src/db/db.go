package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client
var courses *mongo.Collection
var users *mongo.Collection
var problems *mongo.Collection
var submissions *mongo.Collection
var Session *mongo.Collection

func InitDB(client *mongo.Client) {
	db = client
	courses = db.Database("CourseDB").Collection("courses")
	users = db.Database("UserDB").Collection("users")
	problems = db.Database("ProblemDB").Collection("problems")
	submissions = db.Database("SubmissionDB").Collection("submissions")
	Session = db.Database("SessionDB").Collection("sessions")
}

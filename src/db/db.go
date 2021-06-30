package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client
var courses *mongo.Collection
var users *mongo.Collection

func InitDB(client *mongo.Client) {
	db = client
	courses = db.Database("CourseDB").Collection("courses")
	users = db.Database("CourseDB").Collection("users")
}

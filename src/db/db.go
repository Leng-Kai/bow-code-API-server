package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client

func InitDB(client *mongo.Client) {
	db = client
}

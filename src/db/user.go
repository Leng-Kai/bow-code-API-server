package db

import (
	"context"
	"time"

	. "github.com/Leng-Kai/bow-code-API-server/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(newUser User) (UserID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := users.InsertOne(ctx, newUser)
	return res.InsertedID.(UserID), err
}

func GetSingleUser(filter Filter, sortby Sortby) (User, error) {
	opts := options.FindOne().SetSort(sortby)
	var result_bson bson.M
	var result User

	err := users.FindOne(context.TODO(), filter, opts).Decode(&result_bson)
	if err != nil {
		return result, err
	}

	bson_marshal, _ := bson.Marshal(result_bson)
	_ = bson.Unmarshal(bson_marshal, &result)
	return result, err
}

func GetSingleUserByID(id string) (User, error) {
	filter := bson.D{{"_id", id}}
	sortby := bson.D{}
	return GetSingleUser(filter, sortby)
}

func GetMultipleUsers(filter Filter, sortby Sortby) ([]User, error) {
	opts := options.Find().SetSort(sortby)
	var results_bson []bson.M
	var result User
	var results []User

	cursor, err := users.Find(context.TODO(), filter, opts)
	if err != nil {
		return results, err
	}

	err = cursor.All(context.TODO(), &results_bson)
	if err != nil {
		return results, err
	}

	for _, bs := range results_bson {
		bson_marshal, _ := bson.Marshal(bs)
		_ = bson.Unmarshal(bson_marshal, &result)
		results = append(results, result)
	}

	return results, err
}

func UpdateUser(filter Filter, update Update, upsert bool) (User, error) {
	opts := options.FindOneAndUpdate().SetUpsert(upsert)
	var updated_bson bson.M
	var updated User

	err := users.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updated_bson)
	if err != nil {
		return updated, err
	}

	bson_marshal, _ := bson.Marshal(updated_bson)
	_ = bson.Unmarshal(bson_marshal, &updated)
	return updated, err
}

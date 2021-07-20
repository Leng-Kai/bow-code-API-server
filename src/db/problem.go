package db

import (
	"context"
	"time"

	. "github.com/Leng-Kai/bow-code-API-server/schemas"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateProblem(newProblem Problem) (ID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := problems.InsertOne(ctx, newProblem)
	return res.InsertedID.(ID), err
}

func GetSingleProblem(filter Filter, sortby Sortby) (Problem, error) {
	opts := options.FindOne().SetSort(sortby)
	var result_bson bson.M
	var result Problem

	err := problems.FindOne(context.TODO(), filter, opts).Decode(&result_bson)
	if err != nil {
		return result, err
	}

	bson_marshal, _ := bson.Marshal(result_bson)
	_ = bson.Unmarshal(bson_marshal, &result)
	return result, err
}

func DeleteProblem(filter Filter, projection Projection) (Problem, error) {
	opts := options.FindOneAndDelete().SetProjection(projection)
	var deleted_bson bson.M
	var deleted Problem

	err := problems.FindOneAndDelete(context.TODO(), filter, opts).Decode(&deleted_bson)
	if err != nil {
		return deleted, err
	}

	bson_marshal, _ := bson.Marshal(deleted_bson)
	_ = bson.Unmarshal(bson_marshal, &deleted)
	return deleted, err
}

func UpdateProblem(filter Filter, update Update, upsert bool) (Problem, error) {
	opts := options.FindOneAndUpdate().SetUpsert(upsert)
	var updated_bson bson.M
	var updated Problem

	err := problems.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updated_bson)
	if err != nil {
		return updated, err
	}

	bson_marshal, _ := bson.Marshal(updated_bson)
	_ = bson.Unmarshal(bson_marshal, &updated)
	return updated, err
}

func ReplaceProblem(filter Filter, replacement Problem) error {
	var updated_bson bson.M
	err := problems.FindOneAndReplace(context.TODO(), filter, replacement).Decode(&updated_bson)
	if err != nil {
		return err
	}
	return err
}

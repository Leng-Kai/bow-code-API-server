package db

import (
	"context"
	"time"

	. "github.com/Leng-Kai/bow-code-API-server/schemas"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateBulletin(newBulletin Bulletin) (ID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := bulletins.InsertOne(ctx, newBulletin)
	return res.InsertedID.(ID), err
}

func GetSingleClassroom(filter Filter, sortby Sortby) (Classroom, error) {
	opts := options.FindOne().SetSort(sortby)
	var result_bson bson.M
	var result Classroom

	err := classrooms.FindOne(context.TODO(), filter, opts).Decode(&result_bson)
	if err != nil {
		return result, err
	}

	bson_marshal, _ := bson.Marshal(result_bson)
	_ = bson.Unmarshal(bson_marshal, &result)
	return result, err
}

func GetMultipleClassrooms(filter Filter, sortby Sortby) ([]Classroom, error) {
	opts := options.Find().SetSort(sortby)
	var results_bson []bson.M
	var results []Classroom

	cursor, err := classrooms.Find(context.TODO(), filter, opts)
	if err != nil {
		return results, err
	}

	if err = cursor.All(context.TODO(), &results_bson); err != nil {
		return results, err
	}
	for _, result_bson := range results_bson {
		var result Classroom
		bson_marshal, _ := bson.Marshal(result_bson)
		_ = bson.Unmarshal(bson_marshal, &result)
		results = append(results, result)
	}

	return results, err
}

func DeleteClassroom(filter Filter, projection Projection) (Classroom, error) {
	opts := options.FindOneAndDelete().SetProjection(projection)
	var deleted_bson bson.M
	var deleted Classroom

	err := classrooms.FindOneAndDelete(context.TODO(), filter, opts).Decode(&deleted_bson)
	if err != nil {
		return deleted, err
	}

	bson_marshal, _ := bson.Marshal(deleted_bson)
	_ = bson.Unmarshal(bson_marshal, &deleted)
	return deleted, err
}

func UpdateClassroom(filter Filter, update Update, upsert bool) (Classroom, error) {
	opts := options.FindOneAndUpdate().SetUpsert(upsert)
	var updated_bson bson.M
	var updated Classroom

	err := classrooms.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updated_bson)
	if err != nil {
		return updated, err
	}

	bson_marshal, _ := bson.Marshal(updated_bson)
	_ = bson.Unmarshal(bson_marshal, &updated)
	return updated, err
}

func ReplaceClassroom(filter Filter, replacement Classroom) error {
	var updated_bson bson.M
	err := classrooms.FindOneAndReplace(context.TODO(), filter, replacement).Decode(&updated_bson)
	if err != nil {
		return err
	}
	return err
}
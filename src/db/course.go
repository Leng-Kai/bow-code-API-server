package db

import (
	"context"
	"time"

	. "github.com/Leng-Kai/bow-code-API-server/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateCourse(newCourse Course) (ID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := courses.InsertOne(ctx, newCourse)
	return res.InsertedID.(ID), err
}

func GetSingleCourse(filter Filter, sortby Sortby) (Course, error) {
	opts := options.FindOne().SetSort(sortby)
	var result_bson bson.M
	var result Course

	err := courses.FindOne(context.TODO(), filter, opts).Decode(&result_bson)
	if err != nil {
		return result, err
	}

	bson_marshal, _ := bson.Marshal(result_bson)
	_ = bson.Unmarshal(bson_marshal, &result)
	return result, err
}

func GetMultipleCourses(filter Filter, sortby Sortby) ([]Course, error) {
	opts := options.Find().SetSort(sortby)
	var results_bson []bson.M
	var result Course
	var results []Course

	cursor, err := courses.Find(context.TODO(), filter, opts)
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

func DeleteCourse(filter Filter, projection Projection) (Course, error) {
	opts := options.FindOneAndDelete().SetProjection(projection)
	var deleted_bson bson.M
	var deleted Course

	err := courses.FindOneAndDelete(context.TODO(), filter, opts).Decode(&deleted_bson)
	if err != nil {
		return deleted, err
	}

	bson_marshal, _ := bson.Marshal(deleted_bson)
	_ = bson.Unmarshal(bson_marshal, &deleted)
	return deleted, err
}

func UpdateCourse(filter Filter, update Update, upsert bool) (Course, error) {
	opts := options.FindOneAndUpdate().SetUpsert(upsert)
	var updated_bson bson.M
	var updated Course

	err := courses.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updated_bson)
	if err != nil {
		return updated, err
	}

	bson_marshal, _ := bson.Marshal(updated_bson)
	_ = bson.Unmarshal(bson_marshal, &updated)
	return updated, err
}

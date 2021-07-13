package db

import (
	"context"
	"time"

	. "github.com/Leng-Kai/bow-code-API-server/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetMultipleCourses(filter Filter, sortby Sortby) ([]Course, interface{}, error) {
	opts := options.Find().SetSort(sortby)
	var results_bson []bson.M
	var tags_result_bson []bson.M
	type tagCntRes struct {
		Id    string `json:"tag" bson:"_id"`
		Count int    `json:"count" bson:"count"`
	}
	var tags_result tagCntRes
	var tags_count []tagCntRes
	var results []Course

	cursor, err := courses.Find(context.TODO(), filter, opts)
	if err != nil {
		return results, tags_count, err
	}

	err = cursor.All(context.TODO(), &results_bson)
	if err != nil {
		return results, tags_count, err
	}
	for _, bs := range results_bson {
		var result Course
		bson_marshal, _ := bson.Marshal(bs)
		_ = bson.Unmarshal(bson_marshal, &result)
		results = append(results, result)
	}
	groupStage := mongo.Pipeline{
		bson.D{{"$match", filter}},
		bson.D{{"$project", bson.D{{"_id", 1}, {"tags", 1}}}},
		bson.D{{"$unwind", bson.D{{"path", "$tags"}}}},
		bson.D{{"$group", bson.D{{"_id", "$tags"}, {"count", bson.D{{"$sum", 1}}}}}},
	}
	cursor, err = courses.Aggregate(context.TODO(), groupStage)
	if err != nil {
		return results, tags_count, err
	}
	err = cursor.All(context.TODO(), &tags_result_bson)
	if err != nil {
		return results, tags_count, err
	}
	for _, bs := range tags_result_bson {
		bson_marshal, _ := bson.Marshal(bs)
		_ = bson.Unmarshal(bson_marshal, &tags_result)
		tags_count = append(tags_count, tags_result)
	}
	return results, tags_count, err
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

func ReplaceCourse(filter Filter, replacement Course) error {
	var updated_bson bson.M
	err := courses.FindOneAndReplace(context.TODO(), filter, replacement).Decode(&updated_bson)
	if err != nil {
		return err
	}
	return err
}

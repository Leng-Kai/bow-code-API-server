package db

import (
	"context"
	"time"

	. "github.com/Leng-Kai/bow-code-API-server/schemas"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClassroomRecord(newClassroomRecord ClassroomRecord) (ID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := classroomrecords.InsertOne(ctx, newClassroomRecord)
	return res.InsertedID.(ID), err
}

func GetSingleClassroomRecord(filter Filter, sortby Sortby) (ClassroomRecord, error) {
	opts := options.FindOne().SetSort(sortby)
	var result_bson bson.M
	var result ClassroomRecord

	err := classroomrecords.FindOne(context.TODO(), filter, opts).Decode(&result_bson)
	if err != nil {
		return result, err
	}

	bson_marshal, _ := bson.Marshal(result_bson)
	_ = bson.Unmarshal(bson_marshal, &result)
	return result, err
}

func DeleteClassroomRecord(filter Filter, projection Projection) (ClassroomRecord, error) {
	opts := options.FindOneAndDelete().SetProjection(projection)
	var deleted_bson bson.M
	var deleted ClassroomRecord

	err := classroomrecords.FindOneAndDelete(context.TODO(), filter, opts).Decode(&deleted_bson)
	if err != nil {
		return deleted, err
	}

	bson_marshal, _ := bson.Marshal(deleted_bson)
	_ = bson.Unmarshal(bson_marshal, &deleted)
	return deleted, err
}

func UpdateClassroomRecord(filter Filter, update Update, upsert bool) (ClassroomRecord, error) {
	opts := options.FindOneAndUpdate().SetUpsert(upsert)
	var updated_bson bson.M
	var updated ClassroomRecord

	err := classroomrecords.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updated_bson)
	if err != nil {
		return updated, err
	}

	bson_marshal, _ := bson.Marshal(updated_bson)
	_ = bson.Unmarshal(bson_marshal, &updated)
	return updated, err
}

func ReplaceClassroomRecord(filter Filter, replacement ClassroomRecord) error {
	var updated_bson bson.M
	err := classroomrecords.FindOneAndReplace(context.TODO(), filter, replacement).Decode(&updated_bson)
	if err != nil {
		return err
	}
	return err
}
package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const EntitleDatabaseName string = "entitle"
const EntitleCollectionName string = "records"

type ProgramEntitle struct {
	Id    uint16
	Year  uint16
	Month uint16
	Day   uint16
}

type Entitle struct {
	Id              uint32
	ProgramEntitles []ProgramEntitle
	Rest            uint16
}

func FindEntitleRecord(id uint32) *Entitle {
	cl := GetCollection(EntitleDatabaseName, EntitleCollectionName)
	if cl == nil {
		return nil
	}
	result := cl.FindOne(context.TODO(), bson.M{"_id": id})
	if result == nil {
		return nil
	}
	a := new(Entitle)
	err := result.Decode(a)
	if err != nil {
		return nil
	} else {
		return a
	}
}

func FindEntitleRecords(start, end uint32) []*Entitle {
	cl := GetCollection(EntitleDatabaseName, EntitleCollectionName)
	if cl == nil {
		return nil
	}
	var result []*Entitle
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"_id": 1})
	cursor, err := cl.Find(context.TODO(), bson.M{"_id": bson.M{"$ge": start, "$lt": end}}, findOptions)
	if err != nil {
		return nil
	}
	defer cursor.Close(context.TODO())
	cursor.All(context.TODO(), &result)
	return result
}

func FindEntitleRecordsLive(start, end uint32) []*Entitle {
	cl := GetCollection(EntitleDatabaseName, EntitleCollectionName)
	if cl == nil {
		return nil
	}
	var result []*Entitle
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"_id": 1})
	cursor, err := cl.Find(
		context.TODO(),
		bson.M{"_id": bson.M{"$ge": start, "$lt": end}, "rest": bson.M{"$gt": 0}},
		findOptions)
	if err != nil {
		return nil
	}
	defer cursor.Close(context.TODO())
	cursor.All(context.TODO(), &result)
	return result
}

func UpdateEntitle(e *Entitle) bool {
	cl := GetCollection(EntitleDatabaseName, EntitleCollectionName)
	if cl == nil {
		return false
	}

	result := cl.FindOne(context.TODO(), bson.M{"_id": e.Id})
	if result == nil {
		_, err := cl.InsertOne(context.TODO(), e)
		return err == nil
	} else {
		_, err := cl.ReplaceOne(context.TODO(), bson.M{"_id": e.Id}, e)
		return err == nil
	}
}

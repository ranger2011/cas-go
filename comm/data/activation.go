package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ActivationDatabaseName string = "activation"
const ActivationCollectionName string = "records"

type Activation struct {
	Id       uint32
	AreaCode uint16
	Rest     uint16
}

func (a *Activation) ID() uint32 {
	return a.Id
}

func FindActivationRecords(start, end uint32) []*Activation {
	cl := GetCollection(ActivationDatabaseName, ActivationCollectionName)
	if cl == nil {
		return nil
	}
	var result []*Activation
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

func FindActivationRecordsLive(start, end uint32) []*Activation {
	cl := GetCollection(ActivationDatabaseName, ActivationCollectionName)
	if cl == nil {
		return nil
	}
	var result []*Activation
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

func FindActivationRecord(id uint32) *Activation {
	cl := GetCollection(ActivationDatabaseName, ActivationCollectionName)
	if cl == nil {
		return nil
	}
	result := cl.FindOne(context.TODO(), bson.M{"_id": id})
	if result == nil {
		return nil
	}
	a := new(Activation)
	err := result.Decode(a)
	if err != nil {
		return nil
	} else {
		return a
	}
}

func UpdateActivation(a *Activation) bool {
	cl := GetCollection(ActivationDatabaseName, ActivationCollectionName)
	if cl == nil {
		return false
	}

	result := cl.FindOne(context.TODO(), bson.M{"_id": a.Id})
	if result == nil {
		_, err := cl.InsertOne(context.TODO(), a)
		return err == nil
	} else {
		_, err := cl.ReplaceOne(context.TODO(), bson.M{"_id": a.Id}, a)
		return err == nil
	}
}

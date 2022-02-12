package data

import (
	"context"
	"sort"

	"github.com/ranger2011/cas-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MaxProductSize int = 1024
const ProductDatabaseName string = "product"
const ProductChannelCollectionName = "channel"
const ProductPackageCollectionName = "package"

type ProductChannel struct {
	Id                uint16
	Name              string
	TransportStreamId uint16
	ServiceId         uint16
}

func (product *ProductChannel) ID() uint16 {
	return product.Id
}

type ProductPackage struct {
	Id             uint16
	Name           string
	ChannelIdArray []uint16
}

func (product *ProductPackage) ID() uint16 {
	return product.Id
}

var channelArray []ProductChannel
var packageArray []ProductPackage

func InitChannelArray() {
	cl := GetCollection(ProductDatabaseName, ProductChannelCollectionName)
	if cl == nil {
		return
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"_id": 1})
	cusor, err := cl.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return
	}
	defer cusor.Close(context.TODO())
	cusor.All(context.TODO(), &channelArray)
}

func InitPackageArray() {
	cl := GetCollection(ProductDatabaseName, ProductPackageCollectionName)
	if cl == nil {
		return
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"_id": 1})
	cusor, err := cl.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return
	}
	defer cusor.Close(context.TODO())
	cusor.All(context.TODO(), &packageArray)
}

func AddChannel(product *ProductChannel) bool {
	cl := GetCollection(ProductDatabaseName, ProductChannelCollectionName)
	if cl == nil {
		return false
	}
	for i := 0; i < len(channelArray); i++ {
		if channelArray[i].Id == product.Id {
			filter := bson.M{"_id": product.Id}
			_, err := cl.ReplaceOne(context.TODO(), filter, product, options.Replace().SetUpsert(true))
			if err != nil {
				return false
			}
			channelArray[i] = *product
			return true
		}
	}
	_, err := cl.InsertOne(context.TODO(), product)
	if err != nil {
		return false
	}
	channelArray = append(channelArray, *product)
	sort.SliceStable(channelArray, func(i, j int) bool {
		return channelArray[i].Id < channelArray[j].Id
	})
	return true
}

func FindChannel(productId uint16) (bool, *ProductChannel) {
	for _, product := range channelArray {
		if product.Id == productId {
			return true, &product
		}
	}
	return false, nil
}

func DeleteChannel(productId uint16) bool {
	for i := 0; i < len(channelArray); i++ {
		if channelArray[i].Id == productId {
			cl := GetCollection(ProductDatabaseName, ProductChannelCollectionName)
			if cl != nil {
				_, err := cl.DeleteOne(context.TODO(), bson.M{"_id": productId})
				if err != nil {
					return false
				}
				channelArray = append(channelArray[:i], channelArray[i+1:]...)
				return true
			}

			return false
		}
	}
	return false
}

func AddPackage(product *ProductPackage) bool {
	cl := GetCollection(ProductDatabaseName, ProductPackageCollectionName)
	if cl != nil {
		return false
	}
	for i := 0; i < len(packageArray); i++ {
		if packageArray[i].Id == product.Id {
			filter := bson.M{"_id": product.Id}
			_, err := cl.ReplaceOne(context.TODO(), filter, product, options.Replace().SetUpsert(true))
			if err != nil {
				return false
			}
			packageArray[i] = *product
			return true
		}
	}
	_, err := cl.InsertOne(context.TODO(), product)
	if err != nil {
		return false
	}
	packageArray = append(packageArray, *product)
	sort.SliceStable(packageArray, func(i, j int) bool {
		return packageArray[i].Id < packageArray[j].Id
	})
	return true
}

func FindPackage(productId uint16) (bool, *ProductPackage) {
	for _, product := range packageArray {
		if product.Id == productId {
			return true, &product
		}
	}
	return false, nil
}

func DeletePackage(productId uint16) bool {
	for i := 0; i < len(packageArray); i++ {
		if packageArray[i].Id == productId {
			cl := GetCollection(ProductDatabaseName, ProductChannelCollectionName)
			if cl != nil {
				_, err := cl.DeleteOne(context.TODO(), bson.M{"_id": productId})
				if err != nil {
					return false
				}
				packageArray = append(packageArray[:i], packageArray[:i+1]...)
				return true
			}
		}
	}
	return false
}

func FindPackagesIncludeChannel(channelId uint16) []uint16 {
	result := make([]uint16, 0)
	for _, product := range packageArray {
		if utils.ItemExists(product.ChannelIdArray, channelId) {
			result = append(result, product.Id)
		}
	}
	return result
}

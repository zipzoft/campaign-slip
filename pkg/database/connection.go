package database

import (
	"campiagn-slip/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

type MongoConnect struct {
	Client   *mongo.Client
	context  context.Context
	Database *mongo.Database
}

func (connector *MongoConnect) Disconnect() {
	err := connector.Client.Disconnect(connector.context)

	if err != nil {
		return
	}
}

func MongoConnection() MongoConnect {
	configuration := config.GetConfig()

	client, err := mongo.NewClient(options.Client().ApplyURI(configuration.DBUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, connectErr := context.WithTimeout(context.Background(), 10*time.Second)
	if connectErr != nil {
		//
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	giantDatabase := client.Database(configuration.DBName)

	fmt.Println("Connect success")

	return MongoConnect{
		context:  ctx,
		Client:   client,
		Database: giantDatabase,
	}
}

func Aggregate(col string, pipeline interface{}, option ...*options.AggregateOptions) (interface{}, error) {
	Mongo := MongoConnection()
	collection := Mongo.Database.Collection(col)
	result, err := collection.Aggregate(Mongo.context, pipeline, option...)
	//defer Mongo.Disconnect()

	return result, err
}

func InsertOne(col string, doc interface{}) (*mongo.InsertOneResult, error) {
	Mongo := MongoConnection()
	collection := Mongo.Database.Collection(col)
	result, err := collection.InsertOne(Mongo.context, doc)
	//defer Mongo.Disconnect()

	return result, err
}

func UpdateOne(col string, filter interface{}, doc interface{}) (*mongo.UpdateResult, error) {
	Mongo := MongoConnection()
	collection := Mongo.Database.Collection(col)
	data := bson.M{
		"$set": doc,
	}
	result, err := collection.UpdateOne(Mongo.context, filter, data)
	//defer Mongo.Disconnect()

	return result, err
}

func Find(col string, query interface{}, model interface{}, option ...*options.FindOptions) (interface{}, error) {
	Mongo := MongoConnection()
	collection := Mongo.Database.Collection(col)
	result, err := collection.Find(Mongo.context, query, option...)

	//defer Mongo.Disconnect()

	if err != nil {
		return nil, err
	}

	if err = result.All(Mongo.context, model); err != nil {
		panic(err)
	}

	return model, err
}

func FindOne(col string, filter map[string]interface{}) *mongo.SingleResult {
	Mongo := MongoConnection()
	collection := Mongo.Database.Collection(col)
	//defer Mongo.Disconnect()

	result := collection.FindOne(Mongo.context, filter)

	return result
}

func Pagination(col string, query map[string]interface{}, model interface{}, pageNumber string, limit string, option ...*options.FindOptions) (interface{}, error) {
	Mongo := MongoConnection()
	collection := Mongo.Database.Collection(col)

	//defer Mongo.Disconnect()

	queryFilter := bson.M{}
	for key, val := range query {
		queryFilter[key] = val
	}

	page, _ := strconv.Atoi(pageNumber)
	perPage, _ := strconv.Atoi(limit)
	total, _ := collection.CountDocuments(Mongo.context, query)

	findOptions := options.Find()

	if limit == "" {
		perPage = 10
	}

	if page == 1 || pageNumber == "" {
		page = 1
		findOptions.SetLimit(int64(perPage))
		findOptions.SetSkip(0)
	} else {
		findOptions.SetLimit(int64(perPage))
		findOptions.SetSkip(int64(page-1) * 10)
	}

	option = append(option, findOptions)

	result, err := collection.Find(Mongo.context, query, option...)
	if err != nil {
		return nil, err
	}

	if err = result.All(Mongo.context, model); err != nil {
		return nil, err
	}

	data := map[string]interface{}{}
	data["page"] = page
	data["total"] = total
	data["data"] = model

	return data, err
}

func FindOneAndUpdate(col string, filter interface{}, update interface{}) *mongo.SingleResult {
	Mongo := MongoConnection()
	collection := Mongo.Database.Collection(col)
	result := collection.FindOneAndUpdate(Mongo.context, filter, update)
	//defer Mongo.Disconnect()

	return result
}

func FindOneAndDelete(col string, filter interface{}) *mongo.SingleResult {
	Mongo := MongoConnection()
	collection := Mongo.Database.Collection(col)
	result := collection.FindOneAndDelete(Mongo.context, filter)
	//defer Mongo.Disconnect()

	return result
}

func CountDocument(col string, filter interface{}) (int64, error) {
	Mongo := MongoConnection()
	collection := Mongo.Database.Collection(col)
	result, err := collection.CountDocuments(Mongo.context, filter)

	if err != nil {
		return 0, err
	}

	return result, nil
}

func InsertMany(col string, docs []interface{}) (*mongo.InsertManyResult, error) {
	Mongo := MongoConnection()
	//defer Mongo.Disconnect()
	collection := Mongo.Database.Collection(col)
	result, err := collection.InsertMany(Mongo.context, docs)
	return result, err
}

package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/nischayjain4948/golang/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Nischay:Nischay@crudd.qoju4.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT CONNECTION
var collection *mongo.Collection

// Connect with mongoDB
func init() {

	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect with mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption) //context.TODO()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection successfully")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(colName))

	// collection Instance
	fmt.Println("Collections Instance is ready")

}

// Write method to handle the  data for dataBase

// 1. Insert  one record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert movie in db with id: ", inserted.InsertedID)
}

// 2 update one record

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count : ", result.ModifiedCount)
}

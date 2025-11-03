package controller

import (
	"context"
	"fmt"
	"log"
	"mongoapi/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	//client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo db connection successful")

	collection := client.Database(dbName).Collection(colName)
	_ = collection // avoid unused variable; collection can be used elsewhere later

	fmt.Println("connection instance is ready")
}

//mongo db helpers

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatalln("failed to insert the data", err)
	}

	fmt.Println("inserted one movie into database", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatalln("failed to convert string id into object id", err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	updated,err := collection.UpdateOne(context.Background(), filter, update)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("moive status updated successfully",updated.ModifiedCount)
}

func deleteOneMovie(movideId string){
	id,err := primitive.ObjectIDFromHex(movideId)
	if err!= nil {
		log.Fatalln(err)
	}

	filter := bson.M{"_id":id}
	deleted,err := collection.DeleteOne(context.Background(),filter)

	if err!=nil{
		log.Fatalln(err)
	}

	fmt.Println("deleted count",deleted.DeletedCount)
}

//get all collection

func getAllCollection() [] primitive.M{ 

	var moives[] primitive.M

	cursor,err := collection.Find(context.Background(),bson.M{})

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(cursor)

	for cursor.Next(context.Background()){
		var movie bson.M
		err := cursor.Decode(&movie)
		if err!=null{
			log.Fatalln(err)
		}

		moives= append(moives, movie)
	}
	defer cursor.Close(context.Background())
	return moives
}


package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://sukuna:Samar@123@cluster0.0qaccuo.mongodb.net/?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"

// important

var collection *mongo.Collection

//connect with mongodb

func init() { //init runs only one time
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	// conntct
	client,err :=mongo.Connect(context.TODO(), clientOption)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("connection done")

	collection =client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("collection reference is ready")

}

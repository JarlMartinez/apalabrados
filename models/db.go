package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

func ConnectDB() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(db_url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err)
	}

	DB = client.Database(db_name)

	// defer func() {
	// 	err = client.Disconnect(ctx)
	// 	if err != nil {
	// 		fmt.Printf("Error disconnecting from mongo DB")
	// 		panic(err)
	// 	}
	// }()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Printf("Error pinging mongo DB")
		panic(err)
	}

	fmt.Println("Connected to db successfully")

}

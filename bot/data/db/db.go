package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("botDB")
	fmt.Println("Connected to db")
	return db
}

var DB = getDB()
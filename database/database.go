package database

import (
	"context"
	"time"
	"os"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

// Ctx | @desc: mongo context interface
var Ctx context.Context

// Cancel | @desc: mongo context cancel function
var Cancel context.CancelFunc

// Client | @desc: mongo client struct
var Client *mongo.Client

// DB | @desc: mongo database struct
var DB *mongo.Database


func Connect() error {
	err := godotenv.Load(".env")
	if err != nil{
	 log.Fatalf("Error loading .env file: %s", err)
	}
	var err2 error
	Ctx, Cancel = context.WithTimeout(context.Background(), 1000*time.Second)
	Client, err2 = mongo.Connect(Ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err2 != nil {
		panic(err2)
	}

	// Connect to mongo database
	DB = Client.Database("Go-API")
	return nil
}
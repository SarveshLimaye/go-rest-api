package main

import (
	"context"
	"time"
	"os"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
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


const mongoUri = "mongodb+srv://sarvesh:sarvesh2002@cluster0.anzgr.mongodb.net/Go-API?retryWrites=true&w=majority"

type Book struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title,omitempty" bson:"title,omitempty"`
	Price string `json:"price,omitempty" bson:"price,omitempty"`
	Author string `json:"author,omitempty" bson:"author,omitempty"`
}

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

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}else{
		log.Println("Connected to MongoDB")
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/api/v1/books", func(c *fiber.Ctx) error {
		var books []Book
		collection := DB.Collection("books")
		cursor, err := collection.Find(Ctx, bson.M{})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		if err := cursor.All(Ctx, &books); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		 
		return c.JSON(books)

	})

	app.Post("/api/v1/books", func(c *fiber.Ctx) error {
		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		collection := DB.Collection("books")
		res, err := collection.InsertOne(Ctx, book)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(res)

	})

	app.Get("/api/v1/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		var book Book
		collection := DB.Collection("books")
		err = collection.FindOne(Ctx, bson.M{"_id": oid}).Decode(&book)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(book)
	})

	app.Put("/api/v1/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		collection := DB.Collection("books")
		_, err = collection.UpdateOne(Ctx, bson.M{"_id": oid}, bson.M{"$set": book})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendString("Book updated successfully")

	})

	app.Delete("/api/v1/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		oid,err := primitive.ObjectIDFromHex(id)
		if(err != nil){
			return c.Status(400).SendString(err.Error())
		}
		collection := DB.Collection("books")
		_,err2 := collection.DeleteOne(Ctx, bson.M{"_id": oid})
		if(err2 != nil){
			return c.Status(500).SendString(err2.Error())
		}
		
		return c.SendString("Book deleted successfully")
	})

	app.Listen(":3000")
}
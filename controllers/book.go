package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/SarveshLimaye/go-rest-api/database"
	"github.com/SarveshLimaye/go-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	collection := database.DB.Collection("books")
	cursor, err := collection.Find(database.Ctx, bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err := cursor.All(database.Ctx, &books); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	 
	return c.JSON(books)

}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	collection := database.DB.Collection("books")
	res, err := collection.InsertOne(database.Ctx, book)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(res)

}

func GetBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	var book models.Book
	collection := database.DB.Collection("books")
	err = collection.FindOne(database.Ctx, bson.M{"_id": oid}).Decode(&book)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	collection := database.DB.Collection("books")
	_, err = collection.UpdateOne(database.Ctx, bson.M{"_id": oid}, bson.M{"$set": book})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("Book updated successfully")

}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	oid,err := primitive.ObjectIDFromHex(id)
	if(err != nil){
		return c.Status(400).SendString(err.Error())
	}
	collection := database.DB.Collection("books")
	_,err2 := collection.DeleteOne(database.Ctx, bson.M{"_id": oid})
	if(err2 != nil){
		return c.Status(500).SendString(err2.Error())
	}
	
	return c.SendString("Book deleted successfully")
}
package models
import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title,omitempty" bson:"title,omitempty"`
	Price string `json:"price,omitempty" bson:"price,omitempty"`
	Author string `json:"author,omitempty" bson:"author,omitempty"`
}
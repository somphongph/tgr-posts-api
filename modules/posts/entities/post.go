package entities

import (
	"tgr-posts-api/modules/shared/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	Detail   string             `json:"detail"`
	ImageUrl string             `json:"imageUrl"`
	PlaceTag string             `json:"placeTag"`

	entities.Entity `bson:",inline"`
}

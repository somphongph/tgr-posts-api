package domains

import (
	"tgr-posts-api/modules/shared/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	Detail   string             `json:"detail"`
	ImageUrl string             `json:"imageUrl"`

	models.Entity `bson:",inline"`
}

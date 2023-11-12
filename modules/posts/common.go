package posts

import (
	"tgr-posts-api/internal/store/cache"
	models "tgr-posts-api/modules/shared/domains"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	store storer
	cache cache.Cached
}

func InitHandler(store storer, cache cache.Cached) *Handler {
	return &Handler{store: store, cache: cache}
}

type Post struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title"`
	Detail   string             `json:"detail"`
	ImageUrl string             `json:"imageUrl"`

	models.Entity `bson:",inline"`
}

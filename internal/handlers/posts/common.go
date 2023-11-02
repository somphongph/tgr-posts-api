package posts

import (
	"tgr-posts-api/internal/store/cache"
	"time"

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
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title"`
	Detail    string             `json:"detail"`
	ImageUrl  string             `json:"imageUrl"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

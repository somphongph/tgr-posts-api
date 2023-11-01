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
	Caption   string             `json:"caption"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type PostRequest struct {
	Title   string `json:"title"`
	Caption string `json:"caption"`
}

type PostResponse struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Caption string `json:"caption"`
}

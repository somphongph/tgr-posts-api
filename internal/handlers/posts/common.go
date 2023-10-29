package posts

import (
	"tgr-posts-api/internal/store/cache"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	cache cache.Cached
}

func NewHandler(cache cache.Cached) *Handler {
	return &Handler{cache: cache}
}

type Post struct {
	Id      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title"`
	Caption string             `json:"caption"`
}

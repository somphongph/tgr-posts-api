package posts

import (
	"tgr-posts-api/internal/store/cache"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type storer interface {
	GetById(string) (Post, error)
	GetAll() ([]Post, error)
	Add(*Post) error
	Update(*Post) error
	Delete(string) error
}

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

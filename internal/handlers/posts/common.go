package posts

import "tgr-posts-api/internal/cache"

type Handler struct {
	cache cache.Cached
}

func NewHandler(cache cache.Cached) *Handler {
	return &Handler{cache: cache}
}

type Post struct {
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	SubModel string `json:"subModel"`
}

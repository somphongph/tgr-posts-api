package posts

import "tgr-posts-api/internal/repositories/cache"

type Handler struct {
	cache cache.Cached
}

func NewHandler(cache cache.Cached) *Handler {
	return &Handler{cache: cache}
}

type Post struct {
	Title   string `json:"title"`
	Caption string `json:"caption"`
}

package handlers

import (
	"tgr-posts-api/modules/posts/repositories"
	"tgr-posts-api/modules/shared/repositories/cache"
)

type Handler struct {
	store repositories.PostRepository
	cache cache.Cached
}

func PostHandler(store repositories.PostRepository, cache cache.Cached) *Handler {
	return &Handler{store: store, cache: cache}
}
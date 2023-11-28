package handlers

import (
	"tgr-posts-api/modules/posts/repositories"
	"tgr-posts-api/modules/shared/repositories/cache"
)

type handler struct {
	store repositories.PostRepository
	cache cache.Cached
}

func PostHandler(store repositories.PostRepository, cache cache.Cached) *handler {
	return &handler{store: store, cache: cache}
}

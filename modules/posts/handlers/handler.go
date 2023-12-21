package handlers

import (
	"tgr-posts-api/configs"
	"tgr-posts-api/modules/posts/repositories"
	"tgr-posts-api/modules/shared/repositories/cache"
)

type handler struct {
	cfg   *configs.Configs
	store repositories.PostRepository
	cache cache.Cached
}

func PostHandler(cfg *configs.Configs, store repositories.PostRepository, cache cache.Cached) *handler {
	return &handler{cfg: cfg, store: store, cache: cache}
}

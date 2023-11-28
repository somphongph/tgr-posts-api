package routers

import (
	"tgr-posts-api/configs"
	"tgr-posts-api/modules/posts/handlers"
	"tgr-posts-api/modules/posts/repositories"
	"tgr-posts-api/modules/shared/repositories/cache"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, cfg *configs.Configs) {
	// API Version
	api := e.Group("/v1")
	cache := cache.InitRedisCache(&cfg.Redis)

	// Posts
	//---------------------------------------------------
	postRepo := repositories.InitPostRepository(&cfg.MongoDB)
	postHandler := handlers.PostHandler(postRepo, cache)
	postApi := api.Group("/posts")
	{
		postApi.GET("/:id", postHandler.GetItemPostHandler)
		postApi.GET("", postHandler.GetListPostHandler)
		postApi.POST("", postHandler.AddPostHandler)
	}

}

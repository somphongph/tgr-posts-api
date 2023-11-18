package router

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

	// Posts
	//---------------------------------------------------
	p := handlers.PostHandler(repositories.InitMongoDBStore(&cfg.MongoDB), cache.InitCache(&cfg.Redis))
	pApi := api.Group("/posts")
	{
		pApi.GET("/:id", p.GetItemPostHandler)
		pApi.GET("", p.GetListPostHandler)
		pApi.POST("", p.AddPostHandler)
	}

}

package router

import (
	"context"
	"tgr-posts-api/configs"
	"tgr-posts-api/modules/posts/handlers"
	"tgr-posts-api/modules/posts/repositories"
	"tgr-posts-api/modules/shared/repositories/cache"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(cfg *configs.Configs) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// API Version
	api := e.Group("/v1")

	// Posts
	//---------------------------------------------------
	p := handlers.PostHandler(repositories.InitMongoDBStore(&cfg.MongoDB), cache.InitCache(&cfg.Redis))
	pApi := api.Group("/posts")
	{
		pApi.GET("/:id", p.GetItemPostHandler)
		pApi.POST("", p.AddPostHandler)
	}

	//---------------------------------------------------
	// Graceful Shutdown
	//---------------------------------------------------
	go func() {
		if err := e.Start(":" + cfg.App.Port); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

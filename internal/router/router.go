package router

import (
	"context"
	"fmt"
	"tgr-posts-api/internal/handlers/posts"
	"tgr-posts-api/internal/store/cache"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func InitRouter(e *echo.Echo) {
	// Config
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// DB
	// db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatal("Connect to database error", err)
	// }
	// defer db.Close()

	api := e.Group("/v1")

	// Validate
	p := posts.NewHandler(cache.InitCache())
	pApi := api.Group("/posts")
	{
		pApi.GET("", p.GetPostHandler)
		pApi.POST("", p.AddPostHandler)
	}

	// Graceful Shutdown
	go func() {
		port := viper.GetString("app.port")
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed { // Start server
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

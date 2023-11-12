package router

import (
	"context"
	"fmt"
	"log"
	"tgr-posts-api/modules/posts/handlers"
	"tgr-posts-api/modules/posts/repositories"
	"tgr-posts-api/modules/shared/repositories/cache"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
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
	// Add Secret
	// viper.SetConfigName(".env")
	// viper.SetConfigType("env")
	// viper.MergeInConfig()

	err2 := godotenv.Load("./configs/.env")
	if err2 != nil {
		log.Fatal("Error loading .env file")
	}

	// API Version
	api := e.Group("/v1")

	// Posts
	//---------------------------------------------------
	p := handlers.PostHandler(repositories.InitMongoDBStore(), cache.InitCache())
	pApi := api.Group("/posts")
	{
		pApi.GET("/:id", p.GetItemPostHandler)
		pApi.POST("", p.AddPostHandler)
	}

	//---------------------------------------------------
	// Graceful Shutdown
	//---------------------------------------------------
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

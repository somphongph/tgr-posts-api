package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"tgr-posts-api/cmd/router"
	"tgr-posts-api/configs"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	// viper
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// godotenv
	err = godotenv.Load("./configs/.env")
	if err != nil {
		panic(fmt.Errorf("error loading .env file"))
	}

	cfg := new(configs.Configs)

	// App
	cfg.App.Port = viper.GetString("app.port")

	// Database
	cfg.MongoDB.Connection = os.Getenv("MONGO_CONNECTION")
	cfg.MongoDB.DbName = os.Getenv("MONGO_DB_NAME")

	// Redis
	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	cfg.Redis.Pass = os.Getenv("REDIS_PASS")
	cfg.Redis.ShortCache, _ = strconv.Atoi(os.Getenv("REDIS_SHORT_CACHE"))
	cfg.Redis.LongCache, _ = strconv.Atoi(os.Getenv("REDIS_LONG_CACHE"))

	// Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Router
	router.InitRouter(e, cfg)

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

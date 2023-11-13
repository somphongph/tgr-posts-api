package main

import (
	"fmt"
	"os"
	"tgr-posts-api/cmd/router"
	"tgr-posts-api/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	// Config
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	cfg := new(configs.Configs)

	// Echo configs
	cfg.App.Port = viper.GetString("app.port")

	// Database Configs
	cfg.MongoDB.Connection = os.Getenv("MONGO_CONNECTION")
	cfg.MongoDB.DbName = os.Getenv("MONGO_DB_NAME")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Router
	router.InitRouter(e)

	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", viper.GetString("app.port"))
}

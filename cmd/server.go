package main

import (
	"fmt"
	"os"
	"tgr-posts-api/cmd/router"
	"tgr-posts-api/configs"

	"github.com/joho/godotenv"
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

	err = godotenv.Load("./configs/.env")
	if err != nil {
		panic(fmt.Errorf("error loading .env file"))
	}

	cfg := new(configs.Configs)

	// Echo configs
	cfg.App.Port = viper.GetString("app.port")

	// Database Configs
	cfg.MongoDB.Connection = os.Getenv("MONGO_CONNECTION")
	cfg.MongoDB.DbName = os.Getenv("MONGO_DB_NAME")

	// Router
	router.InitRouter(cfg)

	fmt.Println("Please use server.go for main file")
}

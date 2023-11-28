package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Configs struct {
	App     App
	MongoDB MongoDB
	Redis   Redis
}

type App struct {
	Port string
}

// Database
type MongoDB struct {
	Connection string
	DbName     string
}

// Redis
type Redis struct {
	Host       string
	Pass       string
	ShortCache int
	LongCache  int
}

func GetConfig() Configs {
	// viper
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %v", err))
	}

	// godotenv
	err = godotenv.Load("./configs/.env")
	if err != nil {
		panic(fmt.Errorf("error loading .env file"))
	}

	return Configs{
		App: App{
			Port: viper.GetString("app.port"),
		},
		MongoDB: MongoDB{
			Connection: os.Getenv("MONGO_CONNECTION"),
			DbName:     os.Getenv("MONGO_DB_NAME"),
		},
		Redis: Redis{
			Host:       os.Getenv("REDIS_HOST"),
			Pass:       os.Getenv("REDIS_PASS"),
			ShortCache: viper.GetInt("redis.short-cache"),
			LongCache:  viper.GetInt("redis.long-cache"),
		},
	}
}

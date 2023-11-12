package main

import (
	"fmt"
	"tgr-posts-api/internal/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Router
	router.InitRouter(e)

	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", viper.GetString("app.port"))
}

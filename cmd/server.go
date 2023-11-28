package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"tgr-posts-api/configs"
	"tgr-posts-api/routers"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	cfg := configs.GetConfig()

	// Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Router
	routers.InitRouter(e, &cfg)

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

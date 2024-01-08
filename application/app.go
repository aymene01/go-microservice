package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(cxt context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(cxt).Err()
	if err != nil {
		return fmt.Errorf("Fail to connect db %w", err)
	}

	fmt.Println("Starting server ...")

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("Failed to run the server: %w", err)
	}

	return nil
}

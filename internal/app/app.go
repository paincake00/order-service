package app

import (
	"net/http"
	"time"

	httpOrder "github.com/paincake00/order-service/internal/delivery/http"
)

type Application struct {
	config       Config
	router       http.Handler
	orderHandler *httpOrder.OrderHandler
}

func New(config Config) *Application {
	app := &Application{}

	app.config = config

	app.orderHandler = &httpOrder.OrderHandler{}

	app.router = app.createRouter()

	return app
}

func (app *Application) Run() error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      app.router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	return srv.ListenAndServe()
}

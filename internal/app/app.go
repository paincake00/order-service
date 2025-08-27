package app

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/paincake00/order-service/internal/cache"
	dbOrder "github.com/paincake00/order-service/internal/db"
	httpOrder "github.com/paincake00/order-service/internal/delivery/http"
	"github.com/paincake00/order-service/internal/domain/service"
	errorsUtil "github.com/paincake00/order-service/internal/errors"
	"go.uber.org/zap"
)

type Application struct {
	config       Config
	logger       *zap.SugaredLogger
	router       http.Handler
	db           *sql.DB
	orderHandler *httpOrder.OrderHandler
	orderService service.OrderService
}

func New(config Config, logger *zap.SugaredLogger) *Application {
	app := &Application{}

	app.config = config
	app.logger = logger

	db, err := dbOrder.New(
		config.db.addr,
		config.db.driver,
		config.db.maxOpenCons,
		config.db.maxIdleCons,
		config.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal("db error: ", err)
	}
	app.db = db

	dbRepo := dbOrder.NewOrderRepository(db)
	cacheRepo := cache.NewLRUCache(config.cache.capacity)
	orderService := service.NewOrderService(dbRepo, cacheRepo)

	app.orderService = orderService

	webErrorService := &errorsUtil.WebErrorService{Logger: logger}

	app.orderHandler = &httpOrder.OrderHandler{
		OrderService: orderService,
		ErrorService: webErrorService,
		Logger:       logger,
	}

	app.router = app.createRouter()

	return app
}

func (app *Application) Run() error {
	if err := app.orderService.RestoreCache(context.Background()); err != nil {
		app.logger.Fatal(err)
	}

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      app.router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.logger.Infof("Signal %s received, shutting down...", s)

		shutdown <- srv.Shutdown(ctx)
	}()

	app.logger.Infof("Service has started successfully")

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	app.logger.Infof("Service is shutting down...")

	err = <-shutdown
	if err != nil {
		app.logger.Errorf("error occured on server shutting down: %s", err.Error())
		return err
	}

	if err := app.db.Close(); err != nil {
		app.logger.Errorf("error occured on db connection close: %s", err.Error())
		return err
	}

	return nil
}

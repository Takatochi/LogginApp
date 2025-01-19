package app

import (
	"LoggingApp/internal/config"
	"LoggingApp/internal/controllers"
	"LoggingApp/internal/database/migrattion"
	"LoggingApp/internal/repository/postgres"
	"LoggingApp/internal/services"
	"LoggingApp/pkg/database"
	"LoggingApp/pkg/handler"
	"LoggingApp/pkg/middleware"
	"LoggingApp/pkg/server"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	config  *config.Config
	handler *handler.Handler
	srv     *server.Server
	store   database.Database
}

func NewApp(config config.Config) *App {

	store := database.NewGormDatabase(config.Connection).GetDB()
	repo := postgres.NewRepository(store)
	services.NewService(repo)
	app := &App{
		config:  &config,
		handler: handler.NewHandler(services.NewService(repo)),
		srv:     new(server.Server),
		store:   database.NewGormDatabase(config.Connection),
	}
	app.configureRouter()
	migrattion.RunMigrations(app.store)
	return app
}

func (app *App) Start() error {

	go func() {
		if err := app.srv.Run(app.config, app.handler); !errors.Is(err, http.ErrServerClosed) {
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// Block the end
	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := app.srv.Stop(ctx); err != nil {

		_ = fmt.Errorf("failed to stop server: %v", err)
	}
	return nil
}

func (app *App) configureRouter() {

	router := app.handler.Routing()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		middleware.Logging(),
	)

	repo := postgres.NewRepository(app.store.GetDB())

	// Реєстрація сервісів
	controller := controllers.NewBaseController(services.NewService(repo))
	controller.AddSingleController(controllers.NewLogController)

	controller.RegisterRoutes(router)
}

package server

import (
	"context"
	"database/sql"
	"db-cp-last/auth"
	"db-cp-last/auth/delivery"
	"db-cp-last/auth/repository/localstorage"
	"db-cp-last/auth/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

type App struct {
	httpServer *http.Server

	authUseCase auth.UseCase
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", viper.GetString("postgres.url"))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewApp() (*App, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	userRepo := localstorage.NewUserLocalStorage()

	// userRepo := postgres.NewUserRepository(db)
	authUseCase := usecase.NewAuthorizer(
		userRepo,
		viper.GetString("auth.hash_salt"),
		[]byte(viper.GetString("auth.signing_key")),
		viper.GetDuration("auth.token_ttl")*time.Second,
	)

	return &App{
		authUseCase: authUseCase,
	}, nil
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Endpoints
	api := router.Group("/auth")
	delivery.RegisterHTTPEndpoints(api, a.authUseCase)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

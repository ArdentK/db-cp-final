package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ArdentK/db-cp-final/auth"
	"github.com/ArdentK/db-cp-final/auth/delivery"
	authsql "github.com/ArdentK/db-cp-final/auth/repository/postgres"
	"github.com/ArdentK/db-cp-final/auth/usecase"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"

	athletcomp "github.com/ArdentK/db-cp-final/pkg/athletComp"
	acDelivery "github.com/ArdentK/db-cp-final/pkg/athletComp/delivery"
	acsql "github.com/ArdentK/db-cp-final/pkg/athletComp/repository/postgres"
	acusecase "github.com/ArdentK/db-cp-final/pkg/athletComp/usecase"

	"github.com/ArdentK/db-cp-final/pkg/competitions"
	compDelivery "github.com/ArdentK/db-cp-final/pkg/competitions/delivery"
	compsql "github.com/ArdentK/db-cp-final/pkg/competitions/repository/postgres"
	compusecase "github.com/ArdentK/db-cp-final/pkg/competitions/usecase"
)

type App struct {
	httpServer *http.Server

	authUseCase auth.UseCase

	compUseCase competitions.CompUseCase

	acUseCase athletcomp.ACUseCase
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

// func NewApp() (*App, error) {
// 	db, err := initDB()
// 	if err != nil {
// 		return nil, err
// 	}
// 	// defer db.Close()

// 	acRepo := acsql.NewAcRepository(db)
// 	acUseCase := acusecase.NewACUseCase(acRepo)

// 	compRepo := compsql.NewCompRepository(db)
// 	compUseCase := compusecase.NewCompUseCase(compRepo)

// 	userRepo := authsql.NewUserRepository(db)
// 	authUseCase := usecase.NewAuthUseCase(
// 		userRepo,
// 		viper.GetString("auth.hash_salt"),
// 		[]byte(viper.GetString("auth.signing_key")),
// 		viper.GetDuration("auth.token_ttl")*time.Second,
// 	)

// 	return &App{
// 		authUseCase: authUseCase,
// 		compUseCase: compUseCase,
// 		acUseCase:   acUseCase,
// 	}, nil
// }

func Run(port string) error {
	db, err := initDB()
	if err != nil {
		return err
	}
	defer db.Close()

	acRepo := acsql.NewAcRepository(db)
	acUseCase := acusecase.NewACUseCase(acRepo)

	compRepo := compsql.NewCompRepository(db)
	compUseCase := compusecase.NewCompUseCase(compRepo)

	userRepo := authsql.NewUserRepository(db)
	authUseCase := usecase.NewAuthUseCase(
		userRepo,
		viper.GetString("auth.hash_salt"),
		[]byte(viper.GetString("auth.signing_key")),
		viper.GetDuration("auth.token_ttl"),
	)

	a := &App{
		authUseCase: authUseCase,
		compUseCase: compUseCase,
		acUseCase:   acUseCase,
	}

	// Init gin handler
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Endpoints
	api := router.Group("/auth")
	delivery.RegisterHTTPEndpoints(api, a.authUseCase)

	authMiddleware := delivery.NewAuthMiddleware(a.authUseCase)
	newApi := router.Group("/api", authMiddleware)

	acDelivery.RegisterHTTPEndpoints(newApi, a.acUseCase)

	compDelivery.RegisterHTTPEndpoints(&router.RouterGroup, a.compUseCase)

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

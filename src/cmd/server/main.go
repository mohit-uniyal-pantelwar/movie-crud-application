package main

import (
	"fmt"
	"movie-crud-application/src/internal/adapters/persistance"
	"movie-crud-application/src/internal/config"
	"movie-crud-application/src/internal/interface/input/api/rest/handler"
	"movie-crud-application/src/internal/interface/input/api/rest/routes"
	"movie-crud-application/src/internal/usecase"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func main() {

	//setup logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	config, err := config.LoadConfig(sugar)
	if err != nil {
		sugar.Errorf("failed to load config: %v", err)
		os.Exit(1)
	}

	database, err := persistance.NewDatabase(config)
	if err != nil {
		sugar.Errorf("failed to connect to databas: %v", err)
		os.Exit(1)
	}

	sugar.Info("Connected to database")

	movieRepo := persistance.NewMovieRepo(database)
	userRepo := persistance.NewUserRepo(database)
	sessionRepo := persistance.NewSessionRepo(database)

	movieService := usecase.NewMovieService(movieRepo)
	userService := usecase.NewUserService(userRepo, sessionRepo, config.JWT_SECRET)

	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(config, userService)

	router := routes.InitRoutes(&movieHandler, &userHandler, config.JWT_SECRET)

	err = http.ListenAndServe(fmt.Sprintf(":%s", config.APP_PORT), router)
	if err != nil {
		sugar.Errorf("failed to start server: %v", err)
		os.Exit(1)
	}
}

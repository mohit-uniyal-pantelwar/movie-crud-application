package main

import (
	"fmt"
	"log"
	"movie-crud-application/src/internal/adapters/persistance"
	"movie-crud-application/src/internal/config"
	"movie-crud-application/src/internal/interface/input/api/rest/handler"
	"movie-crud-application/src/internal/interface/input/api/rest/routes"
	"movie-crud-application/src/internal/usecase"
	"net/http"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	database, err := persistance.NewDatabase(config)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Connected to database")

	movieRepo := persistance.NewMovieRepo(database)
	userRepo := persistance.NewUserRepo(database)
	sessionRepo := persistance.NewSessionRepo(database)

	movieService := usecase.NewMovieService(movieRepo)
	userService := usecase.NewUserService(userRepo, sessionRepo)

	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(config, userService)

	router := routes.InitRoutes(&movieHandler, &userHandler, config)

	err = http.ListenAndServe(fmt.Sprintf(":%s", config.APP_PORT), router)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

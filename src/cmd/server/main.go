package main

import (
	"fmt"
	"log"
	"movie-crud-application/src/internal/adapters/persistance"
	"movie-crud-application/src/internal/config"
	moviehandler "movie-crud-application/src/internal/interface/input/api/rest/handler"
	"movie-crud-application/src/internal/interface/input/api/rest/routes"
	movieservice "movie-crud-application/src/internal/usecase"
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
	movieService := movieservice.NewMovieService(movieRepo)
	movieHandler := moviehandler.NewMovieHandler(movieService)

	router := routes.InitRoutes(&movieHandler)

	err = http.ListenAndServe(fmt.Sprintf(":%s", config.APP_PORT), router)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

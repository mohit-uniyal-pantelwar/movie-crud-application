package routes

import (
	moviehandler "movie-crud-application/src/internal/interface/input/api/rest/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(
	movieHandler *moviehandler.MovieHandler,
) http.Handler {
	router := chi.NewRouter()

	router.Route("/movies", func(r chi.Router) {
		r.Get("/", movieHandler.GetMoviesHandler)
		r.Get("/{id}", movieHandler.GetMovieHandler)
		r.Post("/", movieHandler.InsertMovieHandler)
		r.Delete("/{id}", movieHandler.DeleteMovieHandler)
		r.Put("/", movieHandler.UpdateMovieHandler)
	})

	return router
}

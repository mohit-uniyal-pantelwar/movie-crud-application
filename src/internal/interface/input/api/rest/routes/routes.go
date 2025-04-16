package routes

import (
	"movie-crud-application/src/internal/interface/input/api/rest/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(
	movieHandler *handler.MovieHandlerImpl,
	userHandler *handler.UserHandler,
) http.Handler {
	router := chi.NewRouter()

	router.Route("/movies", func(r chi.Router) {
		r.Get("/", (*movieHandler).GetMoviesHandler)
		r.Get("/{id}", (*movieHandler).GetMovieHandler)
		r.Post("/", (*movieHandler).InsertMovieHandler)
		r.Delete("/{id}", (*movieHandler).DeleteMovieHandler)
		r.Put("/", (*movieHandler).UpdateMovieHandler)
	})

	router.Route("/user", func(r chi.Router) {
		r.Post("/", userHandler.RegisterUserHandler)
	})

	return router
}

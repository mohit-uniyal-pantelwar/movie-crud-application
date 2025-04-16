package routes

import (
	"movie-crud-application/src/internal/config"
	"movie-crud-application/src/internal/interface/input/api/rest/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(
	movieHandler *handler.MovieHandler,
	userHandler *handler.UserHandler,
	config *config.Config,
) http.Handler {
	router := chi.NewRouter()

	router.Route("/movies", func(r chi.Router) {
		r.Get("/", movieHandler.GetMoviesHandler)
		r.Get("/{id}", movieHandler.GetMovieHandler)
		r.Post("/", movieHandler.InsertMovieHandler)
		r.Delete("/{id}", movieHandler.DeleteMovieHandler)
		r.Put("/", movieHandler.UpdateMovieHandler)
	})

	router.Route("/auth", func(r chi.Router) {
		r.Post("/register", userHandler.RegisterUserHandler)
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			userHandler.LoginHandler(w, r, config)
		})
	})

	return router
}

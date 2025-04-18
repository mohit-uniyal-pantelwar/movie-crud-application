package routes

import (
	"movie-crud-application/src/internal/interface/input/api/rest/handler"
	"movie-crud-application/src/internal/interface/input/api/rest/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(
	movieHandler *handler.MovieHandler,
	userHandler *handler.UserHandler,
	jwtKey string,
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
		r.Post("/login", userHandler.LoginHandler)
		r.Post("/refresh", userHandler.RefreshTokenHandler)
	})

	router.Route("/user", func(r chi.Router) {

		r.Use(middleware.Authenticate(jwtKey))
		r.Get("/profile", userHandler.GetProfileHandler)
		r.Post("/logout", userHandler.LogoutHandler)
	})

	return router
}

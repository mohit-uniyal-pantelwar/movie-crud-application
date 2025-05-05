package routes

import (
	"movie-crud-application/src/internal/interface/input/api/rest/handler"
	"movie-crud-application/src/internal/interface/input/api/rest/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func InitRoutes(
	movieHandler *handler.MovieHandler,
	userHandler *handler.UserHandler,
	socketHandler *handler.SocketHandler,
	jwtKey string,
) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

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

	router.Get("/ws", socketHandler.UpgradeConnctionHandler)
	// https://medium.com/wisemonks/implementing-websockets-in-golang-d3e8e219733b
	//redis
	//https://medium.com/better-programming/internals-workings-of-redis-718f5871be84

	return router
}

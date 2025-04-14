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

	return router
}

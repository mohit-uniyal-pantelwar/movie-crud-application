package moviehandler

import movieservice "movie-crud-application/src/internal/usecase"

type MovieHandler struct {
	movieService movieservice.MovieService
}

func NewMovieHandler(usecase movieservice.MovieService) MovieHandler {
	return MovieHandler{
		movieService: usecase,
	}
}

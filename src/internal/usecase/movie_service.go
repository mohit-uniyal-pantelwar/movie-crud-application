package movieservice

import "movie-crud-application/src/internal/adapters/persistance"

type MovieService struct {
	movieRepo persistance.MovieRepo
}

func NewMovieService(movieRepo persistance.MovieRepo) MovieService {
	return MovieService{movieRepo: movieRepo}
}

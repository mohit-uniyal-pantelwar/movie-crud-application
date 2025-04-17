package usecase

import (
	"errors"
	models "movie-crud-application/src/internal/core"
	movie "movie-crud-application/src/internal/core"
)

type MovieServiceImpl interface {
	GetAllMovies() ([]movie.Movie, error)
	GetMovieById(id string) (movie.Movie, error)
	InsertMovie(movie movie.Movie) (*movie.Movie, error)
	DeleteMovieById(id string) error
	UpdateMovie(movie movie.Movie) error
}

type MovieService struct {
	movieRepo models.MovieRepoImpl
}

func NewMovieService(movieRepo models.MovieRepoImpl) MovieServiceImpl {
	return MovieService{movieRepo: movieRepo}
}

func (ms MovieService) GetAllMovies() ([]movie.Movie, error) {
	movies, err := ms.movieRepo.GetAllMovies()
	if err != nil {
		return []movie.Movie{}, errors.New("failed to fetch movies, try again later")
	}

	return movies, nil
}

func (ms MovieService) GetMovieById(id string) (movie.Movie, error) {
	movie, err := ms.movieRepo.GetMovieById(id)
	if err != nil {
		return movie, errors.New("failed to get movie, try again later")
	}

	return movie, nil
}

func (ms MovieService) InsertMovie(movie movie.Movie) (*movie.Movie, error) {
	insertedMovie, err := ms.movieRepo.InsertMovie(movie)
	if err != nil {
		return nil, errors.New("failed to insert movie, try again later")
	}

	return insertedMovie, nil
}

func (ms MovieService) DeleteMovieById(id string) error {
	err := ms.movieRepo.DeleteMovieById(id)
	if err != nil {
		return errors.New("failed to delete movie")
	}

	return nil
}

func (ms MovieService) UpdateMovie(movie movie.Movie) error {
	err := ms.movieRepo.UpdateMovie(movie)
	if err != nil {
		return errors.New("failed to udpate movie, try again later")
	}

	return nil
}

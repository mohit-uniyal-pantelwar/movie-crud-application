package handler

import (
	"encoding/json"
	models "movie-crud-application/src/internal/core"
	"movie-crud-application/src/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MovieHandler struct {
	movieService usecase.MovieService
}

func NewMovieHandler(usecase usecase.MovieService) MovieHandler {
	return MovieHandler{
		movieService: usecase,
	}
}

func (mh MovieHandler) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := mh.movieService.GetAllMovies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func (mh MovieHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	movie, err := mh.movieService.GetMovieById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func (mh MovieHandler) InsertMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	insertedMovie, err := mh.movieService.InsertMovie(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(insertedMovie)
}

func (mh MovieHandler) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := mh.movieService.DeleteMovieById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Movie Deleted Successfully"))
}

func (mh MovieHandler) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err := mh.movieService.UpdateMovie(movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Movie Updated Successfully"))
}

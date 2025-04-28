package handler

import (
	"encoding/json"
	models "movie-crud-application/src/internal/core"
	"movie-crud-application/src/internal/usecase"
	"movie-crud-application/src/pkg"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MovieHandler struct {
	movieService usecase.MovieServiceImpl
}

func NewMovieHandler(usecase usecase.MovieServiceImpl) MovieHandler {
	return MovieHandler{
		movieService: usecase,
	}
}

func (mh MovieHandler) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := mh.movieService.GetAllMovies()
	if err != nil {
		pkg.SetReponse(w, http.StatusInternalServerError, map[string]string{}, err.Error(), map[string]string{})
		return
	}
	pkg.SetReponse(w, http.StatusOK, map[string]string{"Content-Type": "application/json"}, "", movies)
}

func (mh MovieHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	movie, err := mh.movieService.GetMovieById(id)
	if err != nil {
		pkg.SetReponse(w, http.StatusInternalServerError, map[string]string{}, err.Error(), map[string]string{})
		return
	}

	pkg.SetReponse(w, http.StatusOK, map[string]string{"Content-Type": "application/json"}, "", movie)
}

func (mh MovieHandler) InsertMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		pkg.SetReponse(w, http.StatusBadRequest, map[string]string{}, err.Error(), map[string]string{})
		return
	}

	insertedMovie, err := mh.movieService.InsertMovie(movie)
	if err != nil {
		pkg.SetReponse(w, http.StatusInternalServerError, map[string]string{}, err.Error(), map[string]string{})
		return
	}

	pkg.SetReponse(w, http.StatusOK, map[string]string{"Content-Type": "application/json"}, "", insertedMovie)
}

func (mh MovieHandler) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := mh.movieService.DeleteMovieById(id)
	if err != nil {
		pkg.SetReponse(w, http.StatusInternalServerError, map[string]string{}, err.Error(), map[string]string{})
		return
	}

	pkg.SetReponse(w, http.StatusOK, map[string]string{}, "Movie Deleted Successfully", map[string]string{})
}

func (mh MovieHandler) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		pkg.SetReponse(w, http.StatusBadRequest, map[string]string{}, err.Error(), map[string]string{})
		return
	}

	err := mh.movieService.UpdateMovie(movie)
	if err != nil {
		pkg.SetReponse(w, http.StatusInternalServerError, map[string]string{}, err.Error(), map[string]string{})
		return
	}

	pkg.SetReponse(w, http.StatusOK, map[string]string{}, "Movie Updated Successfully", map[string]string{})
}

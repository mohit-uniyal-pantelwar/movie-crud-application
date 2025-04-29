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
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	response := pkg.Response{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
		Headers:        map[string]string{"Content-Type": "application/json"},
		Data:           movies,
	}
	response.Set()

}

func (mh MovieHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	movie, err := mh.movieService.GetMovieById(id)
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	response := pkg.Response{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
		Headers:        map[string]string{"Content-Type": "application/json"},
		Data:           movie,
	}
	response.Set()

}

func (mh MovieHandler) InsertMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusBadRequest,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	insertedMovie, err := mh.movieService.InsertMovie(movie)
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	response := pkg.Response{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
		Headers:        map[string]string{"Content-Type": "application/json"},
		Data:           insertedMovie,
	}
	response.Set()

}

func (mh MovieHandler) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := mh.movieService.DeleteMovieById(id)
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	response := pkg.Response{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
		Headers:        map[string]string{"Content-Type": "application/json"},
		Message:        "Movie Deleted Successfully",
	}
	response.Set()

}

func (mh MovieHandler) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusBadRequest,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	err := mh.movieService.UpdateMovie(movie)
	if err != nil {
		response := pkg.Response{
			ResponseWriter: w,
			StatusCode:     http.StatusInternalServerError,
			Error:          err.Error(),
		}
		response.Set()
		return
	}

	response := pkg.Response{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
		Headers:        map[string]string{"Content-Type": "application/json"},
		Message:        "Movie Updated Successfully",
	}
	response.Set()

}

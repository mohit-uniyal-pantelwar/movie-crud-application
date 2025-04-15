package persistance

import (
	"context"
	movie "movie-crud-application/src/internal/core"
	"time"
)

type MovieRepo struct {
	db *Database
}

func NewMovieRepo(d *Database) MovieRepo {
	return MovieRepo{db: d}
}

func (mr MovieRepo) GetAllMovies() ([]movie.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var movies []movie.Movie
	query := "SELECT * FROM movies"
	rows, err := mr.db.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var movie movie.Movie
		if err := rows.Scan(&movie.Id, &movie.Name, &movie.Genre, &movie.Rating, &movie.LengthInMinutes, &movie.Language); err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return movies, nil

}

func (mr MovieRepo) GetMovieById(id string) (movie.Movie, error) {
	var movie movie.Movie

	query := "SELECT * FROM movies WHERE id=$1"
	err := mr.db.db.QueryRow(query, id).Scan(&movie.Id, &movie.Name, &movie.Genre, &movie.Rating, &movie.LengthInMinutes, &movie.Language)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (mr MovieRepo) InsertMovie(movie movie.Movie) (*movie.Movie, error) {
	var id int
	query := "INSERT INTO movies(name, genre, rating, length_in_minutes, language) VALUES($1, $2, $3, $4, $5) RETURNING id"

	err := mr.db.db.QueryRow(query, movie.Name, movie.Genre, movie.Rating, movie.LengthInMinutes, movie.Language).Scan(&id)
	if err != nil {
		return nil, err
	}

	movie.Id = id
	return &movie, err
}

func (mr MovieRepo) DeleteMovieById(id string) error {
	query := "DELETE FROM movies WHERE id = $1"
	_, err := mr.db.db.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (mr MovieRepo) UpdateMovie(movie movie.Movie) error {
	query := `UPDATE movies SET name=$1, genre=$2, rating=$3, length_in_minutes=$4, language=$5 WHERE id=$6`

	_, err := mr.db.db.Query(query, movie.Name, movie.Genre, movie.Rating, movie.LengthInMinutes, movie.Language, movie.Id)
	if err != nil {
		return err
	}

	return nil
}

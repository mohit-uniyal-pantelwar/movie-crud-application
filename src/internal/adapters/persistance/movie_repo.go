package persistance

type MovieRepo struct {
	db *Database
}

func NewMovieRepo(d *Database) MovieRepo {
	return MovieRepo{db: d}
}

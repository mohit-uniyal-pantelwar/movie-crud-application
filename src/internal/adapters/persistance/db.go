package persistance

import (
	"database/sql"
	"fmt"
	"movie-crud-application/src/internal/config"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(config *config.Config) (*Database, error) {

	dataBaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME, config.DB_SSLMODE)

	db, err := sql.Open("postgres", dataBaseUrl)
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

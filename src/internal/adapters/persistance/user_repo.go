package persistance

import (
	models "movie-crud-application/src/internal/core"
	"movie-crud-application/src/pkg"
)

type UserRepoImpl interface {
	CreateUser(user models.User) (models.User, error)
	FindUserByUsername(username string) (models.User, error)
}

type UserRepo struct {
	db *Database
}

func NewUserRepo(d *Database) UserRepoImpl {
	return UserRepo{db: d}
}

func (ur UserRepo) CreateUser(user models.User) (models.User, error) {
	var id int
	query := "INSERT INTO users(username, name, email, password) VALUES($1, $2, $3, $4) RETURNING id"

	hashedPassword, err := pkg.HashPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}

	err = ur.db.db.QueryRow(query, user.Username, user.Name, user.Email, hashedPassword).Scan(&id)
	if err != nil {
		return models.User{}, err
	}

	user.Id = id

	return user, nil
}

func (ur UserRepo) FindUserByUsername(username string) (models.User, error) {
	var user models.User

	query := "SELECT id, username, name, email, password FROM users WHERE username=$1"
	err := ur.db.db.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

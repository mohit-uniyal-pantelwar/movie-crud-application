package usecase

import (
	"errors"
	"movie-crud-application/src/internal/adapters/persistance"
	models "movie-crud-application/src/internal/core"
)

type UserService struct {
	userRepo persistance.UserRepo
}

func NewUserService(userRepo persistance.UserRepo) UserService {
	return UserService{userRepo: userRepo}
}

func (us UserService) CreateUser(user models.User) (models.User, error) {
	createdUser, err := us.userRepo.CreateUser(user)
	if err != nil {
		return createdUser, errors.New("failed to create user, try again later")
	}

	return createdUser, nil
}

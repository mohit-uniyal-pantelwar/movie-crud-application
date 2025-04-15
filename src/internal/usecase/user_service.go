package usecase

import "movie-crud-application/src/internal/adapters/persistance"

type UserService struct {
	userRepo persistance.UserRepo
}

func NewUserService(userRepo persistance.UserRepo) UserService {
	return UserService{userRepo: userRepo}
}

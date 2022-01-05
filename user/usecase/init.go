package usecase

import (
	repo "github.com/flip-clean/user/repository"
)

type userUsecase struct {
	UserRepo repo.UserRepo
}

// NewUserUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewUserUsecase(userRepo repo.UserRepo) UserUsecase {
	return &userUsecase{
		UserRepo: userRepo,
	}
}

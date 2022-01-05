package usecase

import (
	"context"

	"github.com/flip-clean/models"
)

//go:generate moq -out usecase_mock.go . UserUsecase
type UserUsecase interface {
	RegisterUser(ctx context.Context, param models.UserRegistrationPayload) (token string, err error)
	GetUserInfo(ctx context.Context) (username, token string, err error)
	IsUserExist(username string) (isFound bool)
}

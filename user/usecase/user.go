package usecase

import (
	"context"

	"github.com/flip-clean/models"
	"github.com/flip-clean/pkg"
)

func (u *userUsecase) RegisterUser(ctx context.Context, param models.UserRegistrationPayload) (token string, err error) {

	_, isFound := u.UserRepo.GetUser(param.Username)
	if isFound {
		err = models.ErrUserExist
		return
	} else {
		u.UserRepo.SetUser(param.Username, param.Token)
	}

	return
}

func (u *userUsecase) GetUserInfo(ctx context.Context) (username, token string, err error) {
	username = pkg.GetUsernameFromCtx(ctx)
	token, _ = u.UserRepo.GetUser(username)
	return
}

func (u *userUsecase) IsUserExist(username string) (isFound bool) {
	_, isFound = u.UserRepo.GetUser(username)
	return
}

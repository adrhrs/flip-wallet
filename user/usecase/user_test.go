package usecase

import (
	"context"
	"testing"

	"github.com/flip-clean/models"
	"github.com/flip-clean/user/repository"
)

func Test_userUsecase_RegisterUser(t *testing.T) {
	type args struct {
		ctx   context.Context
		param models.UserRegistrationPayload
	}
	tests := []struct {
		name      string
		u         *userUsecase
		args      args
		wantToken string
		wantErr   bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				param: models.UserRegistrationPayload{
					Username: "adr",
				},
			},
			u: &userUsecase{
				UserRepo: &repository.UserRepoMock{
					SetUserFunc: func(username, token string) {},
					GetUserFunc: func(username string) (string, bool) {
						return "", false
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := tt.u.RegisterUser(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotToken != tt.wantToken {
				t.Errorf("userUsecase.RegisterUser() = %v, want %v", gotToken, tt.wantToken)
			}
		})
	}
}

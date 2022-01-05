package memory

import (
	repo "github.com/flip-clean/user/repository"
)

type userMemoryRepo struct {
	userData map[string]string
}

func NewUserMemoryRepo(userData map[string]string) repo.UserRepo {
	return &userMemoryRepo{
		userData: userData,
	}
}

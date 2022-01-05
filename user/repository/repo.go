package repository

//go:generate moq -out repo_mock.go . UserRepo
type UserRepo interface {
	SetUser(username, token string)
	GetUser(username string) (data string, isFound bool)
}

package delivery

import (
	"net/http"

	"github.com/flip-clean/middleware"
	userUsecase "github.com/flip-clean/user/usecase"
)

type UserHandler struct {
	Usecase userUsecase.UserUsecase
}

func NewUserHandler(usecase userUsecase.UserUsecase) {
	handler := &UserHandler{
		Usecase: usecase,
	}
	http.HandleFunc("/health", handler.getServerHealth)
	http.HandleFunc("/create_user", handler.registerUser)
	http.Handle("/get_user", middleware.AuthUser(handler.getUserInfo))

}

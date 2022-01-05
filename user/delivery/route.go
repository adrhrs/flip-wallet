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
	http.HandleFunc("/api/v1/system/health", handler.getServerHealth)
	http.HandleFunc("/api/v1/user/create_user", handler.registerUser)
	http.Handle("/api/v1/user/get_user", middleware.AuthUser(handler.getUserInfo))

}

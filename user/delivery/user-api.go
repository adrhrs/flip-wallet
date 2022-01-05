package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/flip-clean/models"
)

const (
	healthyServerMsg = "server is healthy"
)

func (u *UserHandler) getServerHealth(w http.ResponseWriter, req *http.Request) {

	var (
		t              = time.Now()
		responseStatus = http.StatusOK

		data models.CommonResponse
	)

	defer func() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		json.NewEncoder(w).Encode(data)
	}()

	data = models.CommonResponse{
		Msg:     healthyServerMsg,
		Latency: time.Since(t).Seconds(),
	}
}

func (u *UserHandler) registerUser(w http.ResponseWriter, req *http.Request) {

	var (
		responseStatus = http.StatusCreated
		ctx            = req.Context()

		userToken               string
		userRegistrationPayload models.UserRegistrationPayload
	)

	defer func() {
		data := models.UserRegistrationResponse{
			Token: userToken,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		json.NewEncoder(w).Encode(data)
	}()

	errDecode := json.NewDecoder(req.Body).Decode(&userRegistrationPayload)
	if errDecode != nil {
		log.Println(errDecode)
		responseStatus = http.StatusBadRequest
		return
	}

	userToken, errAuth := authenticateUser(userRegistrationPayload)
	if errAuth != nil {
		log.Println(errAuth)
		responseStatus = http.StatusUnauthorized
		return
	}
	userRegistrationPayload.Token = userToken

	_, errUserRegistration := u.Usecase.RegisterUser(ctx, userRegistrationPayload)
	if errUserRegistration != nil {
		switch errUserRegistration {
		case models.ErrUserExist:
			responseStatus = http.StatusConflict
		default:
			responseStatus = http.StatusInternalServerError
			log.Println(errUserRegistration)
		}
		return
	}
}

func (u *UserHandler) getUserInfo(w http.ResponseWriter, req *http.Request) {

	var (
		responseStatus = http.StatusCreated
		ctx            = req.Context()

		userToken, username string
	)

	defer func() {
		data := models.UserRegistrationResponse{
			Token:    userToken,
			Username: username,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		json.NewEncoder(w).Encode(data)
	}()

	username, userToken, errUserRegistration := u.Usecase.GetUserInfo(ctx)
	if errUserRegistration != nil {
		switch errUserRegistration {
		case models.ErrUserExist:
			responseStatus = http.StatusConflict
		default:
			responseStatus = http.StatusInternalServerError
			log.Println(errUserRegistration)
		}
		return
	}
}

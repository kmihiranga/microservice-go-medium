package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"authentication-service/api/presenter"
	"authentication-service/entity"
	"authentication-service/usecase/user"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService *user.Service
}

func RegisterUserHandler(r *mux.Router, userService *user.Service) {
	handler := &UserHandler{
		userService: userService,
	}
	r.HandleFunc("/user", handler.createUser).Methods(http.MethodPost, http.MethodOptions)
}

// user create handler
func (handler *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	configureCorsHeaders(w, r, "*", "*")
	if r.Method == http.MethodOptions {
		return
	}

	createUserReq := &presenter.UserCreateRequest{}
	_ = json.NewDecoder(r.Body).Decode(&createUserReq)

	userData := &entity.User{
		Email:     createUserReq.Email,
		FirstName: createUserReq.FirstName,
		LastName:  createUserReq.LastName,
		Password:  createUserReq.Password,
	}

	status, err := handler.userService.CreateUser(ctx, userData)
	if err != nil {
		log.Fatalf("Error parsing user creation response data to JSON. %v", err)
		processResponseErrorStatus(w, err, http.StatusExpectationFailed)
		_, _ = w.Write([]byte("Error processing user creation"))
		return
	}

	responsePayload := &presenter.UserCreateResponse{
		AuthCode: status,
	}

	response, err := json.Marshal(responsePayload)
	if err != nil {
		log.Fatalf("Error parsing user creation response data to JSON. %v", err)
		processResponseErrorStatus(w, err, http.StatusExpectationFailed)
		_, _ = w.Write([]byte("Error processing user creatoion"))
		return
	}

	configureDefaultHeaders(w, r)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

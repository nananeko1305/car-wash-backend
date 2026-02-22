package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/nananeko1305/car-wash-backend/internal/domain"
	"github.com/nananeko1305/car-wash-backend/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func InitUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (handler *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := handler.userService.GetAllUsers(r.Context())
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("error with get users from db"))
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (hanlder *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var request domain.CreateUserRequest
	var err error

	w.Header().Set("Content-Type", "application/json")

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("error with unmarshaling json: \n", err)
		return
	}

	err = hanlder.userService.CreateUser(r.Context(), request)
	if errors.Is(err, domain.ErrEmailAlreadyExists) {
		w.WriteHeader(400)
		w.Write([]byte("email already exist"))
		return
	}

	w.WriteHeader(201)
	w.Write([]byte("Succesfully created user!"))
}

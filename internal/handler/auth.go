package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/nananeko1305/car-wash-backend/internal/domain"
	"github.com/nananeko1305/car-wash-backend/internal/service"
	"github.com/nananeko1305/car-wash-backend/pkg/httputil"
)

type AuthHandler struct {
	userService *service.UserService
}

func InitAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (authHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var loginRequest domain.LoginUserRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error with parsing json %s", err)
		return
	}

	generatedToken, err := authHandler.userService.Login(r.Context(), loginRequest)
	if err != nil {
		if errors.Is(err, domain.ErrIncorrectPassword) {
			httputil.WriteError(w, 401, "incorrect password")
			return
		}
		if errors.Is(err, domain.ErrUserNotFound) {
			httputil.WriteError(w, 404, "user not found")
			return
		}
		httputil.WriteError(w, 500, "internal server error")
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"token": generatedToken})
}

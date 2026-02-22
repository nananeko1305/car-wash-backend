package httputil

import (
	"encoding/json"
	"net/http"

	"github.com/nananeko1305/car-wash-backend/internal/domain"
)

func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(domain.AppError{Code: status, Message: message})
}

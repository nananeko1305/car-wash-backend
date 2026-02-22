package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/nananeko1305/car-wash-backend/internal/domain"
	"github.com/nananeko1305/car-wash-backend/internal/jwt"
	"github.com/nananeko1305/car-wash-backend/pkg/httputil"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("Authorization")

		if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
			httputil.WriteError(w, 401, "authorization header missing")
			return
		}

		tokenSplit := strings.Split(authorization, " ")
		jwtToken := tokenSplit[1]

		claims, err := jwt.VerifyJwtToken(jwtToken)
		if errors.Is(err, domain.ErrInvalidToken) {
			httputil.WriteError(w, 401, "invalid token")
			return
		}

		// Put user id in context to use it later in app
		ctx := context.WithValue(r.Context(), UserIDKey, claims.ID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

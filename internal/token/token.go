package token

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nananeko1305/car-wash-backend/internal/db"
)

func GenerateJWTToken(user *db.User) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SECURITY"))

	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    user.Email.String,
		ID:        strconv.Itoa(int(user.ID)),
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := generatedToken.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

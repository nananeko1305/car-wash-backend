package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nananeko1305/car-wash-backend/internal/db"
	"github.com/nananeko1305/car-wash-backend/internal/domain"
)

func GenerateJwtToken(user *db.User) (string, error) {
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

func VerifyJwtToken(jwtToken string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECURITY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil || !token.Valid {
		return nil, domain.ErrInvalidToken
	}
	return claims, nil

}

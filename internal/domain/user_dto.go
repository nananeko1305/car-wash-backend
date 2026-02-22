package domain

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nananeko1305/car-wash-backend/internal/db"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r CreateUserRequest) ToCreateUserParams() db.CreateUserParams {

	var createUserParams db.CreateUserParams

	createUserParams.Name = pgtype.Text{String: r.Name, Valid: true}
	createUserParams.LastName = pgtype.Text{String: r.LastName, Valid: true}
	createUserParams.Email = pgtype.Text{String: r.Email, Valid: true}
	createUserParams.Password = pgtype.Text{String: r.Password, Valid: true}

	return createUserParams
}

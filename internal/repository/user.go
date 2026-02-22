package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nananeko1305/car-wash-backend/internal/db"
	"github.com/nananeko1305/car-wash-backend/internal/domain"
)

type UserRepository struct {
	queries *db.Queries
}

func InitUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{
		queries: queries,
	}
}

func (userRepository *UserRepository) GetAllUsers(ctx context.Context) ([]db.User, error) {
	users, err := userRepository.queries.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (userRepository *UserRepository) CreateUser(ctx context.Context, user db.CreateUserParams) error {
	var pgErr *pgconn.PgError
	err := userRepository.queries.CreateUser(ctx, user)
	if errors.As(err, &pgErr) && pgErr.ConstraintName == "users_email_unique" {
		fmt.Println(pgErr.ConstraintName)
		return domain.ErrEmailAlreadyExists
	}
	return err
}

func (userRepository *UserRepository) GetUserByEmail(ctx context.Context, email pgtype.Text) (*db.User, error) {
	user, err := userRepository.queries.GetUserByEmail(ctx, email)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

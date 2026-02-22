package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nananeko1305/car-wash-backend/internal/db"
	"github.com/nananeko1305/car-wash-backend/internal/domain"
	"github.com/nananeko1305/car-wash-backend/internal/repository"
	"github.com/nananeko1305/car-wash-backend/internal/token"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func InitUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) GetAllUsers(ctx context.Context) ([]db.User, error) {
	users, err := userService.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (userService *UserService) CreateUser(ctx context.Context, request domain.CreateUserRequest) error {

	user := request.ToCreateUserParams()

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = pgtype.Text{String: string(password), Valid: true}

	err = userService.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (userService *UserService) Login(ctx context.Context, request domain.LoginUserRequest) (string, error) {

	email := pgtype.Text{String: request.Email, Valid: true}

	user, err := userService.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(request.Password))
	if err != nil {
		return "", domain.ErrIncorrectPassword
	}

	generatedToken, err := token.GenerateJWTToken(user)
	if err != nil {
		return "", err
	}

	return generatedToken, nil
}

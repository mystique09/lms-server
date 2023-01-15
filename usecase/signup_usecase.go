package usecase

import (
	"server/database/postgresql"
	"server/domain"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type signupUsecase struct {
	userRepository domain.UserRepository
}

func NewSignupUsecase(userRepository domain.UserRepository) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
	}
}

func (su *signupUsecase) GetUserByUsername(c echo.Context, username string) (domain.User, error) {
	return su.userRepository.GetByUsername(c.Request().Context(), username)
}

func (su *signupUsecase) CreateUser(c echo.Context, username, email, password string) error {
	return su.userRepository.Create(c.Request().Context(), &postgresql.CreateUserParams{
		ID:         uuid.New(),
		Username:   username,
		Email:      email,
		Password:   password,
		UserRole:   postgresql.RoleSTUDENT,
		Visibility: postgresql.VisibilityPUBLIC,
	})
}

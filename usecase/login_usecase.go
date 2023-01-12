package usecase

import (
	"server/domain"
	"server/internal/tokenutil"
	"time"

	"github.com/labstack/echo/v4"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	tokenMaker     tokenutil.Maker
}

func NewLoginUsecase(userRepository domain.UserRepository, tokenMaker tokenutil.Maker) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		tokenMaker:     tokenMaker,
	}
}

func (lu *loginUsecase) GetUserByUsername(c echo.Context, username string) (domain.User, error) {
	return lu.userRepository.GetByUsername(c, username)
}

func (lu *loginUsecase) CreateAccessToken(username string, duration time.Duration) (string, *tokenutil.Payload, error) {
	return lu.tokenMaker.CreateToken(username, duration)
}

func (lu *loginUsecase) CreateRefreshToken(username string, duration time.Duration) (string, *tokenutil.Payload, error) {
	return lu.tokenMaker.CreateToken(username, duration)
}

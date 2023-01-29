package usecase

import (
	"server/domain"
	"server/internal/tokenutil"
	"time"

	"github.com/google/uuid"
)

type refreshTokenUsecase struct {
	tokenMaker tokenutil.Maker
}

func NewRefreshTokenUsecase(tokenMaker tokenutil.Maker) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		tokenMaker: tokenMaker,
	}
}

func (rt *refreshTokenUsecase) ValidateRefreshToken(refreshToken string) (*tokenutil.Payload, error) {
	return rt.tokenMaker.VerifyToken(refreshToken)
}

func (rt *refreshTokenUsecase) CreateAccessToken(username string, uid uuid.UUID, duration time.Duration) (string, *tokenutil.Payload, error) {
	return rt.tokenMaker.CreateToken(username, uid, duration)
}

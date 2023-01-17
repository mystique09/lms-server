package usecase

import (
	"server/domain"
	"server/internal/tokenutil"
)

type accessTokenUsecase struct {
	tokenMaker tokenutil.Maker
}

func NewAccessTokenUsecase(tokenMaker tokenutil.Maker) domain.AccessTokenUsecase {
	return &accessTokenUsecase{
		tokenMaker: tokenMaker,
	}
}

func (at *accessTokenUsecase) ValidateAccessToken(accessToken string) (*tokenutil.Payload, error) {
	return at.tokenMaker.VerifyToken(accessToken)
}

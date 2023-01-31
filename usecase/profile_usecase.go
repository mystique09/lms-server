package usecase

import (
	"context"
	"server/domain"

	uuid "github.com/google/uuid"
)

type profileUsecase struct {
	userRepository domain.UserRepository
}

func NewProfileUsecase(userRepo domain.UserRepository) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepo,
	}
}

func (pfu *profileUsecase) GetProfile(c context.Context, id uuid.UUID) (domain.User, error) {
	return pfu.userRepository.GetByID(c, id)
}

func (pfu *profileUsecase) GetClassrooms(c context.Context, id uuid.UUID) ([]domain.Classroom, error) {
	return pfu.userRepository.GetAllJoinedClassrooms(c, id)
}

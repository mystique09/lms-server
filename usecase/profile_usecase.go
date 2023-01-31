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

func (usruc *profileUsecase) GetProfile(c context.Context, id uuid.UUID) (domain.User, error) {
	// TODO!: fetch the profile table, instead of user table
	return usruc.userRepository.GetByID(c, id)
}

func (usruc *profileUsecase) GetClassrooms(c context.Context, id uuid.UUID) ([]domain.Classroom, error) {
	return usruc.userRepository.GetAllJoinedClassrooms(c, id)
}

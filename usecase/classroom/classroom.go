package usecase

import (
	"server/domain/classroom"
)

type (
	getClassroomUsecase struct {
		repository domain.ClassroomRepository
	}
	createClassroomUsecase struct {
		repository domain.ClassroomRepository
	}
	updateClassroomUsecase struct {
		repository domain.ClassroomRepository
	}

	deleteClassroomUsecase struct {
		repository domain.ClassroomRepository
	}
)

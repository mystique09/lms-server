package controller

import (
	"database/sql"
	"server/domain"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (pfc *ProfileController) GetProfile(c echo.Context) error {
	var id = c.Param("id")
	user_id, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	profile, err := pfc.ProfileUsecase.GetProfile(c.Request().Context(), user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_ONE, profile))
}

func (pfc *ProfileController) GetClassrooms(c echo.Context) error {
	var id = c.Param("id")
	user_id, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	_, err = pfc.ProfileUsecase.GetProfile(c.Request().Context(), user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	classrooms, err := pfc.ProfileUsecase.GetClassrooms(c.Request().Context(), user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_ONE, classrooms))
}

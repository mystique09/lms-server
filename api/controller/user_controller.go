package controller

import (
	"database/sql"
	"server/domain"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (pfc *UserController) GetUser(c echo.Context) error {
	var id = c.Param("id")
	user_id, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	profile, err := pfc.UserUsecase.GetProfile(c.Request().Context(), user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_ONE, profile))
}

func (pfc *UserController) GetProfile(c echo.Context) error {
	var id = c.Param("id")
	user_id, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	profile, err := pfc.UserUsecase.GetProfile(c.Request().Context(), user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_ONE, profile))
}

func (pfc *UserController) GetClassrooms(c echo.Context) error {
	var id = c.Param("id")
	user_id, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	_, err = pfc.UserUsecase.GetProfile(c.Request().Context(), user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	classrooms, err := pfc.UserUsecase.GetClassrooms(c.Request().Context(), user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_ALL, classrooms))
}

package routes

import (
	"fmt"
	"net/http"
	database "server/database/sqlc"
	"server/utils"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) getAllClassworks(c echo.Context) error {
	cid := c.Param("id")
	cuid, err := uuid.Parse(cid)

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	page := c.QueryParam("page")
	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}
	if offset < 0 {
		return c.JSON(400, utils.NewResponse(nil, "page must not be megative"))
	}

	check_classrooms, err := s.DB.GetClass(c.Request().Context(), cuid)
	if err != nil || check_classrooms.ID == uuid.Nil {
		return c.JSON(400, utils.NewResponse(nil, fmt.Sprintf("classroom with id [%v] doesn't exist", cuid)))
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	if jwt_payload.ID != check_classrooms.AdminID {
		return c.JSON(http.StatusUnauthorized, utils.NewResponse(nil, "unauthorized access"))
	}

	classworks, err := s.DB.ListClassworkAdmin(c.Request().Context(), database.ListClassworkAdminParams{
		ClassID: cuid,
		Offset:  int32(offset * 10),
	})

	return c.JSON(200, utils.NewResponse(classworks, ""))
}

func (s *Server) getAllUserClassworks(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	if offset < 0 {
		return c.JSON(400, utils.NewResponse(nil, "offset must not be a negative"))
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	if uid != jwt_payload.ID {
		return c.JSON(403, utils.NewResponse(nil, "unauthorized access"))
	}

	classworks, err := s.DB.ListSubmittedClassworks(c.Request().Context(), database.ListSubmittedClassworksParams{
		UserID: uid,
		Offset: int32(offset * 10),
	})

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(200, utils.NewResponse(classworks, ""))
}

func (s *Server) getClassworkById(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	cwid := c.Param("classwork_id")
	parsed_cwid, err := uuid.Parse(cwid)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)
	if err != nil || check_user.ID == uuid.Nil {
		return c.JSON(404, utils.NewResponse(nil, fmt.Sprintf("user with id [%v] doesn't exist", uid)))
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	if jwt_payload.ID != uid {
		return c.JSON(http.StatusUnauthorized, utils.NewResponse(nil, "unauthorized access"))
	}

	classwork, err := s.DB.GetClassWork(c.Request().Context(), database.GetClassWorkParams{
		UserID: uid,
		ID:     parsed_cwid,
	})
	if err != nil || classwork.ID == uuid.Nil {
		return c.JSON(404, utils.NewResponse(nil, fmt.Sprintf("classwork with id [%v] doesn't exist", parsed_cwid)))
	}

	return c.JSON(200, utils.NewResponse(classwork, ""))
}

func (s *Server) addNewClasswork(c echo.Context) error {
	return c.String(200, "TODO: cloud storage for file upload")
}

func (s *Server) deleteClasswork(c echo.Context) error {
	return c.String(200, "TODO")
}

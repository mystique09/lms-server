package routes

import (
	"fmt"
	"net/http"
	database "server/database/sqlc"
	"server/utils"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (s *Server) getAllClassworks(c echo.Context) error {
	cid := c.Param("id")
	cuid, err := uuid.Parse(cid)

	if err != nil {
		return c.JSON(400, err)
	}

	page := c.QueryParam("page")
	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(400, err)
	}
	if offset < 0 {
		return c.JSON(400, NEGATIVE_OFFSET)
	}

	check_classrooms, err := s.DB.GetClass(c.Request().Context(), cuid)
	if err != nil || check_classrooms.ID == uuid.Nil {
		return c.JSON(400, CLASSROOM_NOTFOUND)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	if jwt_payload.ID != check_classrooms.AdminID {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	classworks, err := s.DB.ListClassworkAdmin(c.Request().Context(), database.ListClassworkAdminParams{
		ClassID: cuid,
		Offset:  int32(offset * 10),
	})

	return c.JSON(200, classworks)
}

func (s *Server) getAllUserClassworks(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, err)
	}

	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(400, err)
	}

	if offset < 0 {
		return c.JSON(400, NEGATIVE_OFFSET)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	if uid != jwt_payload.ID {
		return c.JSON(403, UNAUTHORIZED)
	}

	classworks, err := s.DB.ListSubmittedClassworks(c.Request().Context(), database.ListSubmittedClassworksParams{
		UserID: uid,
		Offset: int32(offset * 10),
	})

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, classworks)
}

func (s *Server) getClassworkById(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, err)
	}

	cwid := c.Param("classwork_id")
	parsed_cwid, err := uuid.Parse(cwid)
	if err != nil {
		return c.JSON(400, err)
	}

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)
	if err != nil || check_user.ID == uuid.Nil {
		return c.JSON(404, USER_NOTFOUND)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	if jwt_payload.ID != uid {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	classwork, err := s.DB.GetClassWork(c.Request().Context(), database.GetClassWorkParams{
		UserID: uid,
		ID:     parsed_cwid,
	})
	if err != nil || classwork.ID == uuid.Nil {
		return c.JSON(404, CLASSWORK_NOTFOUND)
	}

	return c.JSON(200, classwork)
}

func (s *Server) addNewClasswork(c echo.Context) error {
	id := c.Param("id")
	cid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, UNAUTHORIZED)
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(400, err)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	check_classrooms, err := s.DB.GetClass(c.Request().Context(), cid)
	if err != nil || check_classrooms.ID == uuid.Nil {
		return c.JSON(400, CLASSROOM_NOTFOUND)
	}

	check_member, err := s.DB.GetClassroomMemberById(c.Request().Context(), database.GetClassroomMemberByIdParams{
		UserID:  jwt_payload.ID,
		ClassID: cid,
	})
	if err != nil || check_member.UserID == uuid.Nil {
		return c.JSON(404, NOT_A_MEMBER)
	}

	file_id := uuid.New()

	src, err := file.Open()
	if err != nil {
		return c.JSON(401, err)
	}
	defer src.Close()

	resp, err := s.Cld.Upload.Upload(c.Request().Context(), src, uploader.UploadParams{
		PublicID:       fmt.Sprintf("class-management/classworks/%v", file_id.String()),
		Transformation: "c_crop,g_center,/q_auto/f_auto",
		Tags:           []string{"assignments", "classworks", file_id.String()},
	})

	new_classwork, err := s.DB.InsertNewClasswork(c.Request().Context(), database.InsertNewClassworkParams{
		ID:      file_id,
		ClassID: cid,
		UserID:  jwt_payload.ID,
		Url:     resp.SecureURL,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, UNAUTHORIZED)
	}

	return c.JSON(200, new_classwork)
}

func (s *Server) deleteClasswork(c echo.Context) error {
	class_id := c.Param("id")
	classword_id := c.Param("classword_id")

	class_uuid, err := uuid.Parse(class_id)
	if err != nil {
		return c.JSON(400, err)
	}

	cw_uuid, err := uuid.Parse(classword_id)
	if err != nil {
		return c.JSON(400, err)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payoad := utils.GetPayloadFromJwt(jwt_token)

	deleted_cw, err := s.DB.DeleteClassworkFromClass(c.Request().Context(), database.DeleteClassworkFromClassParams{
		ClassID: class_uuid,
		ID:      cw_uuid,
		UserID:  jwt_payoad.ID,
	})

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, deleted_cw)
}

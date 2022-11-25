package api

import "github.com/labstack/echo/v4"

func (s *Server) getOneComment(c echo.Context) error {

	return c.String(200, "TODO")
}

func (s *Server) createNewComment(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) updateComment(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) deleteComment(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) getAllCommentLikes(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) likeComment(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) unlikeComment(c echo.Context) error {
	return c.String(200, "TODO")
}

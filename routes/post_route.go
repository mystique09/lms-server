package routes

import "github.com/labstack/echo/v4"

func (s *Server) getOnePost(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) createNewPost(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) updatePost(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) deletePost(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) getAllPostLikes(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) likePost(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) unlikePost(c echo.Context) error {
	return c.String(200, "TODO")
}

package routes

import "github.com/labstack/echo/v4"

func (s *Server) getFollowers(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) getFollowings(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) addNewFollowing(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) removeFollowing(c echo.Context) error {
	return c.String(200, "TODO")
}

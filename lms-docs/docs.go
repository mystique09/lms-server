package lmsdocs

import (
	"embed"
	"io/fs"

	"github.com/labstack/echo/v4"
)

//go:embed swagger
var lmsDocs embed.FS

func BuildStaticHTTPS() fs.FS {
	s := echo.MustSubFS(lmsDocs, "swagger")
	return s
}

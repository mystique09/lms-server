package web

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:dist
var BuildFs embed.FS

func BuildWebFS() fs.FS {
	build, err := fs.Sub(BuildFs, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return build
}

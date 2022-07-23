package frontend

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:build/*
var BuildFs embed.FS

func BuildHTTPFS() fs.FS {
	build, err := fs.Sub(BuildFs, "build")

	if err != nil {
		log.Fatal(err)
	}
	return build
}

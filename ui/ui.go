package ui

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//go:embed  public/*
var UiEmbed embed.FS

func NewTemplateRenderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseFS(UiEmbed, "public/*.html")),
	}
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
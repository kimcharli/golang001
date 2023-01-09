package webserver

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// type Template struct {
// 	templates *template.Template
// }

type TemlateRenderer struct {
	templates *template.Template
}

// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return t.templates.ExecuteTemplate(w, name, data)
// }

func (t *TemlateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "Worlld")
}

func WebServer() (e *echo.Echo) {
	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("public/views/*.html")),
	// }
	renderer := &TemlateRenderer{
		// templates: template.Must(template.ParseGlob("*.html")),
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e = echo.New()

	// e.Renderer = t
	e.Renderer = renderer

	e.GET("/hello", Hello)
	e.GET("/something", func(c echo.Context) error {
		return c.Render(http.StatusOK, "template.html", map[string]interface{}{
			"name": "Dolly",
		})
	}).Name = "foobar"
	return e
}

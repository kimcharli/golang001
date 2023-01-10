package webserver

import (
	"errors"
	"html/template"
	"io"
	"net/http"

	// "github.com/labstack/echo"
	"github.com/labstack/echo/v4"

	"github.com/kimcharli/go101/handler"
)

type TemlateRenderer struct {
	templates map[string]*template.Template
}

// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return t.templates.ExecuteTemplate(w, name, data)
// }

func (t *TemlateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	// if viewContext, isMap := data.(map[string]interface{}); isMap {
	// 	viewContext["reverse"] = c.Echo().Reverse
	// }

	// return t.templates.ExecuteTemplate(w, name, data)
	return tmpl.ExecuteTemplate(w, "layout.html", data)
}

// func exampleStringOut(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello World!")
// }

// func exampleRenderOut(c echo.Context) error {
// 	return c.Render(http.StatusOK, "hello", "W0rld")
// }

// func exampleJSONOut(c echo.Context) error {
// 	return c.JSONBlob(
// 		http.StatusOK,
// 		[]byte(`{"id": "1", "msg": "Hello, Boatswain!"}`),
// 	)
// }

// func exampleHTMLOut(c echo.Context) error {
// 	return c.HTML(
// 		http.StatusOK,
// 		"<h1>Hello, Boatswain!</h1>",
// 	)
// }

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "Worlld")
}

// func SomeThing(c echo.Context) error {

// }

func WebServer() (e *echo.Echo) {

	e = echo.New()

	// e.Renderer = t
	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("view/home.html", "view/layout.html"))
	templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/layout.html"))
	e.Renderer = &TemlateRenderer{
		// templates: template.Must(template.ParseGlob("*.html")),
		// templates: template.Must(template.ParseGlob("view/*.html")),
		templates: templates,
	}

	// e.GET("/hello", Hello)
	// e.GET("/something", func(c echo.Context) error {
	// 	return c.Render(http.StatusOK, "template.html", map[string]interface{}{
	// 		"name": "Dolly",
	// 	})
	// }).Name = "foobar"

	e.GET("/", handler.HomeHander)
	e.GET("/about", handler.AboutHandler)
	e.POST("/clicked", handler.ClickedHandler)
	return e
}

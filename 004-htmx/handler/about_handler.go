package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AboutHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"name": "About",
		"msg":  "All about Charlie",
	})
}

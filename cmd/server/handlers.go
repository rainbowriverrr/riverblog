package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"Name": "Elon Musk",
	})
}

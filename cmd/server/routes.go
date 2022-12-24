package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

// The routes
var homeRoute echo.Route = echo.Route{
	Method: http.MethodGet,
	Path:   "/",
	Handler: func(c echo.Context) error {
		return c.File("templates/mainpage.html")
	},
}

// initializes the routes for the pocketbase app
func initRoutes(e *core.ServeEvent) error {

	e.Router.AddRoute(homeRoute)
	e.Router.Static("/", "frontend/public")

	return nil
}

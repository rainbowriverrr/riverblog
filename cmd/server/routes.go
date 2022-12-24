package main

import (
	"github.com/pocketbase/pocketbase/core"
)

// initializes the routes for the pocketbase app
func initRoutes(e *core.ServeEvent) error {

	e.Router.GET("/", HomeHandler)
	e.Router.Static("/", "frontend/dist")

	return nil
}

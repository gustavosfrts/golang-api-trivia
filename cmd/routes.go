package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavosfrts/goapi-docker/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
}
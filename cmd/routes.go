package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustavosfrts/golang-api-trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Post("/fact", CreateFact)
}
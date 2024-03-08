package server

import "github.com/gofiber/fiber/v2"

func routes(s *Server, app *fiber.App) (routes fiber.Router) {
	routes = app.Group("/v1")
	routes.Post("/journal", s.controller.CreateJournal)
	return
}

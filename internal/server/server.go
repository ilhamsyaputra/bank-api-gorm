package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/controller"
)

type Server struct {
	controller controller.Controller
}

func InitServer(controller controller.Controller) *Server {
	return &Server{
		controller: controller,
	}
}

func (s *Server) Start() {
	app := fiber.New()

	routes := app.Group("/v1")
	routes.Post("/daftar", s.controller.Daftar)
	routes.Post("/tabung", s.controller.Tabung)

	err := app.Listen(":2525")
	if err != nil {
		panic(err)
	}
}

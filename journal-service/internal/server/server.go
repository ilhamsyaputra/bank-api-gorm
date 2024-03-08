package server

import (
	"journal-service/internal/controller"
	"journal-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	controller controller.Controller
}

func InitServer(controller *controller.Controller) *Server {
	return &Server{
		controller: *controller,
	}
}

func (s *Server) Start(logger *logger.Logger) {
	app := fiber.New()
	routes(s, app)

	err := app.Listen(":20026")
	if err != nil {
		panic(err)
	}
}

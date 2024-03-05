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
	routes.Post("/tarik", s.controller.Tarik)
	routes.Post("/transfer", s.controller.Transfer)
	routes.Get("/saldo/:no_rekening", s.controller.CekSaldo)

	err := app.Listen(":2525")
	if err != nil {
		panic(err)
	}
}

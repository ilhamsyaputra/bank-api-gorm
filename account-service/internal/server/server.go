package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/controller"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/middleware"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
)

type Server struct {
	controller controller.Controller
}

func InitServer(controller controller.Controller) *Server {
	return &Server{
		controller: controller,
	}
}

func (s *Server) Start(logger *logger.Logger) {
	app := fiber.New()

	authentication := middleware.NewAuthenticationMiddleware()
	authorization := middleware.NewAuthorizationMiddleware()

	routes := app.Group("/v1")
	routes.Post("/daftar", s.controller.Daftar)
	routes.Post("/login", s.controller.Login)

	routes.Use(authentication)
	routes.Post("/tabung", authorization, s.controller.Tabung)
	routes.Post("/tarik", authorization, s.controller.Tarik)
	routes.Post("/transfer", authorization, s.controller.Transfer)
	routes.Get("/saldo/:no_rekening", authorization, s.controller.CekSaldo)
	routes.Get("/mutasi/:no_rekening", authorization, s.controller.GetMutasi)

	err := app.Listen(":20025")
	if err != nil {
		panic(err)
	}
}

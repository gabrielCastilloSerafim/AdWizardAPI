package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Server struct {
	App *fiber.App
}

func createServer() *Server {
	return &Server{
		App: fiber.New(),
	}
}

func (s *Server) addMiddlewares() {

	s.App.Use(logger.New())
	s.App.Use(requestid.New())
}

func (s *Server) start() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(s.App.Listen("0.0.0.0:" + port))
}

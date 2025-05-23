package server

import (
	"database/sql"
	"os"

	"github.com/Gambitier/grandchattutorial/internal/handlers"
	"github.com/Gambitier/grandchattutorial/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app     *fiber.App
	queries *repository.Queries
}

func New(db *sql.DB) *Server {
	queries := repository.New(db)

	app := fiber.New(fiber.Config{
		AppName: "Grand Chat API",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	return &Server{
		app:     app,
		queries: queries,
	}
}

func (s *Server) SetupRoutes() {
	// Routes
	api := s.app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	handlers.RegisterAuthRoutes(auth, s.queries)

	// Chat routes
	chat := api.Group("/chat")
	chat.Get("/rooms", s.handleGetRooms)
	chat.Post("/rooms", s.handleCreateRoom)
	chat.Get("/rooms/:id", s.handleGetRoom)
	chat.Post("/rooms/:id/messages", s.handleCreateMessage)
	chat.Get("/rooms/:id/messages", s.handleGetMessages)
	chat.Post("/rooms/:id/join", s.handleJoinRoom)
}

func (s *Server) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return s.app.Listen(":" + port)
}

// Handler functions
func (s *Server) handleGetRooms(c *fiber.Ctx) error {
	return c.SendString("Get rooms endpoint")
}

func (s *Server) handleCreateRoom(c *fiber.Ctx) error {
	return c.SendString("Create room endpoint")
}

func (s *Server) handleGetRoom(c *fiber.Ctx) error {
	return c.SendString("Get room endpoint")
}

func (s *Server) handleCreateMessage(c *fiber.Ctx) error {
	return c.SendString("Create message endpoint")
}

func (s *Server) handleGetMessages(c *fiber.Ctx) error {
	return c.SendString("Get messages endpoint")
}

func (s *Server) handleJoinRoom(c *fiber.Ctx) error {
	return c.SendString("Join room endpoint")
}

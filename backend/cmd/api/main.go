package main

import (
	"log"
	"morng-dev/internal/adapters/http/handlers"
	"morng-dev/internal/adapters/http/middleware"
	"morng-dev/internal/adapters/http/routes"
	"morng-dev/internal/adapters/persistence/repositories"
	"morng-dev/internal/config"
	"morng-dev/internal/core/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//loadconfig
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	//db
	db := config.SetupDatabase(cfg)
	//repo
	userRepo := repositories.NewUserRepository(db)
	todoRepo := repositories.NewTodoRepository(db)
	//service
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	todoService := services.NewTodoRepository(todoRepo)
	//middleware
	authMW := middleware.NewAuthMiddleware(cfg.JWT_SECRET)
	//handler
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	todoHandler := handlers.NewTodoHandler(todoService)

	//app fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	//routes
	routes := routes.NewRoutes(
		authHandler,
		todoHandler,
		userHandler,
		authMW,
	)
	routes.SetupRoutes(app)

	log.Printf("Server starting on port %s", cfg.APPPORT)
	log.Fatal(app.Listen(":" + cfg.APPPORT))
}

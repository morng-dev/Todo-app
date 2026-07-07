package routes

import (
	"morng-dev/internal/adapters/http/handlers"
	"morng-dev/internal/adapters/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Routes struct {
	AuthHandler *handlers.AuthHandler
	TodoHandler *handlers.TodoHandler
	UserHandler *handlers.UserHandler
	authMW      *middleware.AuthMiddleware
}

func NewRoutes(
	AuthHandler *handlers.AuthHandler,
	TodoHandler *handlers.TodoHandler,
	UserHandler *handlers.UserHandler,
	authMW *middleware.AuthMiddleware,
) *Routes {
	return &Routes{
		AuthHandler: AuthHandler,
		TodoHandler: TodoHandler,
		UserHandler: UserHandler,
		authMW:      authMW,
	}
}

func (r *Routes) SetupRoutes(app *fiber.App) {
	//middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	//Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "fiber todo is running",
		})
	})
	api := app.Group("/api/v1")

	//auth api
	auth := api.Group("/auth")
	auth.Post("/register", r.AuthHandler.Register)
	auth.Post("/login", r.AuthHandler.Login)

	user := api.Group("/user")
	user.Get("/", r.UserHandler.GetUsers)

	todo := api.Group("/todo")
	todo.Get("/:id", r.TodoHandler.GetTodoByID)
	todo.Get("/", r.TodoHandler.GetAllTodo)

	todoAdmin := todo.Group("", r.authMW.Authrequire())
	todoAdmin.Post("/", r.TodoHandler.CreateTodo)
	todoAdmin.Put("/:id", r.TodoHandler.UpdateTodoStatus)
	todoAdmin.Delete("/:id", r.TodoHandler.DeleteTodo)
}

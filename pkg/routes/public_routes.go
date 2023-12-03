package routes

import (
	"github.com/gofiber/fiber/v2"
	"main/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")
	// Routes for method:
	route.Post("/login", controllers.UserSignIn)
	route.Post("/sign-up", controllers.UserSignUp)
}

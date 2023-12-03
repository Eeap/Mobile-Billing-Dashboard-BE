package routes

import (
	"github.com/gofiber/fiber/v2"
	"main/app/controllers"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for method:
	route.Post("/logout", controllers.UserSignOut)
	route.Get("/aws-resource", controllers.GetResourcesCost)
	route.Get("/alert-messages", controllers.GetAlertMessages)
	route.Post("/user-key", controllers.UserKeySet)
	route.Post("/alert-setting", controllers.AlertSetting)
}

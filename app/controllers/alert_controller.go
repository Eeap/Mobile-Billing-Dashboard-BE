package controllers

import (
	"github.com/gofiber/fiber/v2"
	"main/app/models"
	"main/app/services"
)

// @Router /v1/alert-messages [get]
func GetAlertMessages(c *fiber.Ctx) error {
	email := c.Query("email")
	resData, err := services.GetAlertMessages(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"totalResults": len(*resData),
		"messages":     *resData,
	})
}

// @Router /v1/alert-setting [post]
func AlertSetting(c *fiber.Ctx) error {
	alertSettings := &models.AlertSetting{}
	if err := c.BodyParser(alertSettings); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err,
		})
	}
	res, err := services.UpdateItem(alertSettings)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   res,
	})
}

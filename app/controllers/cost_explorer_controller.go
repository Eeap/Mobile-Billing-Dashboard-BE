package controllers

import (
	"github.com/gofiber/fiber/v2"
	"main/app/models"
	"main/app/services"
)

// @Router /v1/aws-resource [get]
func GetResourcesCost(c *fiber.Ctx) error {
	costIn := &models.CostExplorer{}
	if err := c.QueryParser(costIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}
	resData, err := services.GetCostUsageByResource(costIn)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"totalResults": len(resData),
		"resources":    resData,
	})
}

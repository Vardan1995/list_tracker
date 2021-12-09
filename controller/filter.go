package controller

import (
	"github.com/Vardan1995/filter_notifyer/entity"
	"github.com/Vardan1995/filter_notifyer/service"
	"github.com/gofiber/fiber/v2"
)

var (
	filterService = service.NewFilterService()
)

func PushTracker(c *fiber.Ctx) error {
	id := c.Locals("id").(uint)
	var filter entity.Filter
	filter.UserId = id
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Enter required data correctly!"})
	}

	if _, err := filterService.CreateFilter(&filter); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create filter", "data": err})
	}

	// fmt.Printf("filter: %v\n", filter)
	return c.JSON(fiber.Map{"success": "true"})
}

func GetUserFilters(c *fiber.Ctx) error {
	id := c.Locals("id").(uint)
	var filters []entity.Filter
	err := filterService.GetUserFilter(&filters, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "filter with specified user does not exist"})

	}
	return c.JSON(fiber.Map{"token_data": filters, "success": "true"})
}
func GetAllFilters(c *fiber.Ctx) error {
	var filters []entity.Filter
	err := filterService.GetFilters(&filters)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "filter with specified user does not exist"})

	}
	return c.JSON(fiber.Map{"token_data": filters, "success": "true"})
}

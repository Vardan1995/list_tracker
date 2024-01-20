package controller

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/Vardan1995/list_tracker/entity"
	"github.com/Vardan1995/list_tracker/service"
	"github.com/gofiber/fiber/v3"
)

var (
	preferenceService = service.NewPreferenceService()
)

func PushTracker(c fiber.Ctx) error {
	// var id uint = 1 //c.Locals("id").(uint)
	var preference = entity.Preference{UserId: 1, Link: c.FormValue("link")}

	preference.UserId = 1
	if err := json.Unmarshal(c.Request().Body(), &preference); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Enter required data correctly!"})
	}

	u, err := url.Parse(preference.Link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Maliformed url!"})

	}

	fmt.Printf("Host: %v\n", u.Host) //TODO validate host` list.am/...`

	if _, err := preferenceService.CreatePreference(&preference); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create preference", "data": err})
	}

	// fmt.Printf("preference: %v\n", preference)
	return c.JSON(fiber.Map{"success": "true"})
}

func GetUserPreferences(c fiber.Ctx) error {
	var id uint = 1 //c.Locals("id").(uint)
	var preferences []entity.Preference
	err := preferenceService.GetUserPreference(&preferences, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "preference with specified user does not exist"})

	}
	return c.JSON(fiber.Map{"token_data": preferences, "success": "true"})
}
func GetAllPreferences(c fiber.Ctx) error {
	var preferences []entity.Preference
	err := preferenceService.GetPreferences(&preferences)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "preference with specified user does not exist"})

	}
	return c.JSON(fiber.Map{"token_data": preferences, "success": "true"})
}

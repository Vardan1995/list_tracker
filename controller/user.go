package controller

import (
	"encoding/json"

	"github.com/Vardan1995/list_tracker/entity"
	"github.com/Vardan1995/list_tracker/service"
	"github.com/gofiber/fiber/v3"
)

var (
	userService = service.NewUserService()
)

func CreateUser(c fiber.Ctx) error {
	user := entity.User{}
	if err := json.Unmarshal(c.Request().Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Enter required data correctly!"})
	}

	_, token, err := userService.CreateUser(&user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	return c.JSON(fiber.Map{"token": token})
}

func GetUser(c fiber.Ctx) error {
	id := c.Locals("id").(uint)
	user := entity.User{}

	_, err := userService.GetUser(&user, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User with specified id does not exist"})

	}
	return c.JSON(fiber.Map{"user": user, "success": "true"})
}

func Login(c fiber.Ctx) error {
	type UserData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	var ud UserData
	// err := c.BodyParser(&ud)
	user := entity.User{}
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "can`t parse body"})
	// }

	token, err := userService.LoginUser(&user, ud.Username, ud.Email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"messager": "not found"})
	}

	return c.JSON(fiber.Map{"user": user, "token": token})
}

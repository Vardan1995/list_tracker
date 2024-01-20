package router

import (
	"github.com/Vardan1995/list_tracker/controller"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	auth := api.Group("/auth")
	user := api.Group("/user") //todo auth

	//auth routes
	auth.Post("/login", controller.Login)
	auth.Post("/create", controller.CreateUser)

	//user routes
	user.Get("/", controller.GetUser)
	user.Post("/preference", controller.PushTracker)
	user.Get("/preference", controller.GetUserPreferences)
	user.Get("/preference/all", controller.GetAllPreferences)

}

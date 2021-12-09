package main

import (
	"github.com/Vardan1995/filter_notifyer/core"
	"github.com/Vardan1995/filter_notifyer/database"
	"github.com/Vardan1995/filter_notifyer/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	app.Static("/", "./public")
	database.ConnectDB()
	// Default config
	app.Use(cors.New())
	go core.Track()
	app.Listen("127.0.0.1:5000")
	//----------------------------------------------------------------

}

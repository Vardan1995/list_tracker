package main

import (
	"github.com/Vardan1995/list_tracker/core"
	"github.com/Vardan1995/list_tracker/database"
	"github.com/Vardan1995/list_tracker/router"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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

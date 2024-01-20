package main

import (
	"github.com/Vardan1995/list_tracker/core"
	"github.com/Vardan1995/list_tracker/database"
	"github.com/Vardan1995/list_tracker/router"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	app.Static("/", "./public")
	database.ConnectDB()
	// Default config
	app.Use(cors.New())
	app.Use(logger.New(logger.ConfigDefault))
	go core.Track()
	app.Listen("127.0.0.1:5000")
	//----------------------------------------------------------------

}

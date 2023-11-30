package main

import (
	"github.com/NoIdeaCoder/Saniama/internals/database"
	"github.com/NoIdeaCoder/Saniama/internals/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/django/v3"
)

func main() {
	database.InitUserDatabase()
	database.InitProductDatabase()
	database.InitOrderDatabase()
	engine := django.New("web/templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Static("/static", "web/static")
	routes.HandleRoutes(app)
	app.Listen(":8080")
}

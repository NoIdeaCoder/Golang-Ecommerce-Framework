package routes

import (
	"fmt"

	"github.com/NoIdeaCoder/Saniama/internals/controllers"
	"github.com/NoIdeaCoder/Saniama/internals/helpers"
	"github.com/gofiber/fiber/v2"
)

func HandleRoutes(app *fiber.App) {
	app.Post("/api/signup", func(c *fiber.Ctx) error {
		return controllers.Signup(c)
	})
	app.Post("/api/login", func(c *fiber.Ctx) error {
		return controllers.Login(c)
	})

	app.Get("/api/placeorder/:products", func(c *fiber.Ctx) error {
		return helpers.PlaceOrder(c, c.Params("products"))
	})

	app.Post("/api/addToCart/:productname", func(c *fiber.Ctx) error {
		return helpers.AddToCart(c, c.Params("productname"))
	})
	app.Get("/snippets/:snippet_name", func(c *fiber.Ctx) error {
		if c.Params("snippet_name") == "products" {
			return c.Render(fmt.Sprintf("snippets/%s", c.Params("snippet_name")),
				fiber.Map{
					"products": helpers.GetProducts(),
				})
		}
		if c.Params("snippet_name") == "cart" {
			usercart, err := helpers.GetUserInfo(c)
			if err != nil {
				return c.SendString("Internal Server Error. We're Sorry")
			}
			return c.Render(fmt.Sprintf("snippets/%s", c.Params("snippet_name")),
				fiber.Map{"cart": usercart.Cart},
			)
		}
		return c.Render(fmt.Sprintf("snippets/%s", c.Params("snippet_name")), nil)
	})
	app.Get("/products/:product", func(c *fiber.Ctx) error {
		return c.Render("snippets/product", fiber.Map{
			"product": helpers.GetSpecificProduct(c.Params("product")),
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", nil)
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", nil)
	})

	app.Get("/", controllers.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.Render("home", nil)
	})
}

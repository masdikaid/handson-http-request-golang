package httphandler

import (
	"hands-on-gohttp/product"

	"github.com/gofiber/fiber/v2"
)

func App() {
	app := fiber.New()
	app.Get("/product", func(c *fiber.Ctx) error {
		return c.JSON(product.ProductList())
	})
	app.Listen(":3000")
}

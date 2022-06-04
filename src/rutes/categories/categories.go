package categories

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var socketDB *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	socketDB = sqldb
	app.Get("/api/categories", categories)
	app.Get("/api/sub_categories", sub_categories)
	app.Get("/api/all_categories", all_categories)
}

func categories(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func sub_categories(c *fiber.Ctx) error {
	return c.SendString("sub_categories")
}
func all_categories(c *fiber.Ctx) error {
	return c.SendString("all_categories")
}

package curso

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var socketDB *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	socketDB = sqldb
	app.Get("/api/cursos/top_courses", top_courses)
	app.Get("/api/cursos/category_wise_course", category_wise_course)
	app.Get("/api/cursos/courses_by_search_string", courses_by_search_string)
	app.Get("/api/cursos/filter_course", filter_course)
	app.Get("/api/cursos/my_wishlist", my_wishlist)
	app.Get("/api/cursos/toggle_wishlist_items", toggle_wishlist_items)
	app.Get("/api/cursos/course_details_by_id", course_details_by_id)
	app.Get("/api/cursos/enroll_free_course", enroll_free_course)
}

func top_courses(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func category_wise_course(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func courses_by_search_string(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func filter_course(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func my_wishlist(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func toggle_wishlist_items(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func course_details_by_id(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func enroll_free_course(c *fiber.Ctx) error {
	return c.SendString("Login")
}

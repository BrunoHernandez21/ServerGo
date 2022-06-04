package my_curses

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var socketDB *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	socketDB = sqldb
	app.Get("/api/auth/sections", sections)
	app.Get("/api/auth/save_course_progress", save_course_progress)
	app.Get("/api/auth/enroll_free_course", enroll_free_course)
	app.Get("/api/auth/my_courses", my_courses)
}

// Funciones
func sections(c *fiber.Ctx) error {
	return c.SendString("{\nsections\n}")
}
func save_course_progress(c *fiber.Ctx) error {
	return c.SendString("save_course_progress")
}
func enroll_free_course(c *fiber.Ctx) error {
	return c.SendString("enroll_free_course")
}
func my_courses(c *fiber.Ctx) error {
	return c.SendString("my_courses")
}

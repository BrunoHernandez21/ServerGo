package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init_routes(app *fiber.App, sqldb *gorm.DB) {
	db = sqldb
	v1 := app.Group("/api/auth/", midelwareAuth1, midelwareAuth2)
	v2 := app.Group("/api/auth/")
	v2.Post("signup", signup)
	v1.Get("login", login)
	v1.Get("logout", logout)
	v1.Get("userdata", userdata)
	v1.Get("update_password", update_password)
	v1.Get("update_userdata", update_userdata)
	v1.Get("upload_user_image", upload_user_image)

}

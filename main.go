package main

import (
	"cursos_app/src/config"
	"cursos_app/src/rutes/auth"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//declaracion del servidor Fiber
	app := fiber.New(fiber.Config{
		AppName: "Aldo Nery",
		Prefork: true,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	//Inicio de la base de datos
	db := conexionDB()
	//Rutas
	auth.Init_routes(app, db)
	//Iniciar Servidor
	err2 := app.Listen(":3000")
	if err2 != nil {
		panic(err2.Error())
	}

}

func conexionDB() (conexiones *gorm.DB) {
	dns := config.DB.User + ":" + config.DB.Password + config.DB.Soc + config.DB.TableName
	conexion, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

/*
	go mod init
	go mod tidy
	go get -u gorm.io/gorm
*/

package auth

import (
	"cursos_app/src/models"
	"cursos_app/src/moduls"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

///
func midelwareAuth1(c *fiber.Ctx) error {
	if true {
		return c.Next()
	} else {
		return errors.New("no login")
	}
}
func midelwareAuth2(c *fiber.Ctx) error {
	if true {
		return c.Next()
	} else {
		return errors.New("no login")
	}

}

func login(c *fiber.Ctx) error {

	return c.JSON("{algo,true,true}")
}
func logout(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func signup(c *fiber.Ctx) error {
	// Declare a new struct.
	// input := models.Users_IO{}
	// if err := c.BodyParser(&input); err != nil {
	// 	return err
	// }
	// inputdb := input.To_user()
	// fmt.Print(inputdb)
	//Crea una entrada en la tabla
	//db.AutoMigrate(&inputb)
	//db.Create(&inputdb)
	moduls.Send_Mail_Password("bruno.hgallegos@alumnos.udg.mx")

	//db.Model(&input).Updates(input)
	return c.JSON("")
}

func userdata(c *fiber.Ctx) error {
	input := models.Users{}
	db.First(&input, "id = ?", 2)

	inputdb := input.To_user_io()
	fmt.Print(inputdb)

	return c.JSON(inputdb)
}
func update_password(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func update_userdata(c *fiber.Ctx) error {
	return c.SendString("Login")
}
func upload_user_image(c *fiber.Ctx) error {
	return c.SendString("Login")
}

////////////////// crear tablas
///// crear
//db.Omit("Name", "Age", "CreatedAt").Create(&user)
//db.Select("Name", "Age", "CreatedAt").Create(&user)
// los crea omitiendo o seleccionando

////////////////// Busquedas
/////////////////////////////// Individuales
// db.First(&user)
// // SELECT * FROM users ORDER BY id LIMIT 1;

// // Obtener un registro, sin orden especificado
//  db.Take(&user)
// // SELECT * FROM users LIMIT 1;

// // Obtenga el último registro, ordenado por clave principal desc
//  db.Last(&user)
// // SELECT * FROM users ORDER BY id DESC LIMIT 1;

// db.First(&usuario, "id = ?" , "D04" )
// // Get first matched record
// db.Where("name = ?", "jinzhu").First(&user)
// // SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

/////////////////////////////// Busquedas sencillas
// // Get all matched records
// db.Where("name <> ?", "jinzhu").Find(&users)
// SELECT * FROM users WHERE name <> 'jinzhu';
// // IN
// db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// // SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

// // LIKE
// db.Where("name LIKE ?", "%jin%").Find(&users)
// // SELECT * FROM users WHERE name LIKE '%jin%';

// // AND
// db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
// // SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

// // Time
// db.Where("updated_at > ?", lastWeek).Find(&users)
// // SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// // BETWEEN
// db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// // SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

//////// update
/////////////////////////////////// single
///// este guarda y sobrescribe todo
// db.Save(&user)

//// Actualzaciones de una columna
// Update with conditions
// db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// // UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;
// // User's ID is `111`:
// db.Model(&user).Update("name", "hello")
// // UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;
// // Update with conditions and model value
// db.Model(&user).Where("active = ?", true).Update("name", "hello")
// // UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

/////////////////////////////////// pluri
///Actualiza todo lo no nul value
// // Update attributes with `struct`, will only update non-zero fields
// db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
// // UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

///////// Delete
// db.Delete(&User{}, 10)
// // DELETE FROM users WHERE id = 10;

// db.Delete(&User{}, "10")
// // DELETE FROM users WHERE id = 10;

// db.Delete(&users, []int{1,2,3})
// // DELETE FROM users WHERE id IN (1,2,3);

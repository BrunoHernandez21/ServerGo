package auth

import (
	"crypto/sha1"
	"cursos_app/src/models"
	"cursos_app/src/moduls"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func signup(c *fiber.Ctx) error {

	input := models.Users_IO{}
	if err := c.BodyParser(&input); err != nil {
		return err
	}
	inputdb := input.To_user()

	password := utils.UUID()[0:13]
	h := sha1.New()
	h.Write([]byte(password))
	i := hex.EncodeToString(h.Sum(nil))
	inputdb.Password = i

	moduls.Send_Mail_Password(inputdb.Email, password)

	db.Create(&inputdb)
	return c.JSON(input)
}

package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	models "zendx.io/P2P-Drive/models"
)

var userLogin models.RegisterRequest

// -------------------------- Login User --------------------------\\

func UserLogin(c *fiber.Ctx) error {

	json.Unmarshal(c.Body(), &userLogin)

	//fmt.Print(user)

	//return c.JSON(user)

	Database := Connection()

	token := Database.DBemailCheck(userLogin.Email)

	if token == "Not Found" {
		Database.DBregister(&userLogin)
		return c.JSON(userLogin)

	} else {
		return c.SendString("Account with email exists")
	}
}

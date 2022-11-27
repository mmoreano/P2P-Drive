package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	models "zendx.io/P2P-Drive/models"
)

var user models.RegisterRequest

//-------------------------- Register User --------------------------\\

func UserRegister(c *fiber.Ctx) error {

	json.Unmarshal(c.Body(), &user)

	//fmt.Print(user)

	//return c.JSON(user)

	Database := Connection()

	val := Database.DBemailCheck(user.Email)

	fmt.Print(val)

	if val == "Not Found" {
		Database.DBregister(user.Username, user.UserPassword, user.Number,
			user.Email, user.FirstName, user.LastName)

		return c.JSON(user)

		//u, err := json.Marshal(user)

		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println(u)
		//fmt.Println(string(u))
		//return c.SendString(string(u))

	} else {
		return c.SendString("Account with email exists")
	}
}

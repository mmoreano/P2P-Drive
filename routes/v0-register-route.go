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
	//Creating connection to DB
	Database := Connection()

	//Getting form data from query Params

	user.Username = c.Query("username")
	user.UserPassword = c.Query("password")
	user.Number = c.Query("number")
	user.Email = c.Query("email")
	user.FirstName = c.Query("firstname")
	user.LastName = c.Query("lastname")

	fmt.Println(user.Username)

	//Checking for existing user

	val := Database.DBemailCheck(user.Email)

	// If user exists, return error
	if val == "Not Found" {
		Database.DBregister(user.UserPassword, user.Username, user.Email, user.Number, user.FirstName, user.LastName)

		jsonUser, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}
		fmt.Println(jsonUser)
		fmt.Println(string(jsonUser))
		Database.CloseClientDB()
		return c.SendString(string(jsonUser))
	} else {
		return c.SendString("Account with email exists")
	}
}

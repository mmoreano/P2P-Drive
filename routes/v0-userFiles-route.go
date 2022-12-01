package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	models "zendx.io/P2P-Drive/models"
)

// ------------------------- Get File Status --------------------------\\
var files models.UserFileResponse

func UserFiles(c *fiber.Ctx) error {
	//Fetching query params
	arg := c.Query("Token")
	files := Connection().GetUserFiles(arg)

	jsonData, err := json.Marshal(files)
	if err != nil {
		panic(err)
	}
	return c.SendString(string(jsonData))
}

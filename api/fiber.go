package api

import (
	"encoding/json"
	"fmt"
	"zendx.io/P2P-Drive/routes"
	"github.com/gofiber/fiber/v2"
)

func StartAPI() {
	api := fiber.New()

	api.Post("/push", func(c *fiber.Ctx) error {
		model := routes.Add(c)
		data, _ := json.Marshal(model)
		fmt.Println(string(data))
		return c.SendString(string(data))
	})

	api.Get("/get", func(c *fiber.Ctx) error {
		arg := c.Query("arg")
		data := routes.Get(arg)
		u, err := json.Marshal(data)

		if err != nil {
			panic(err)
		}
		return c.SendString(string(u))
	})



	api.Listen(":8586")
}

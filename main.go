package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"zendx.io/P2P-Drive/setup"
)

func main() {

	log.Print("Starting API...")
	api := fiber.New()
	setup.LaunchEndpoints(api)
	api.Listen(":8083")

}

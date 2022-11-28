package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/swagger"
	"log"
	_ "zendx.io/P2P-Drive/docs"
	"zendx.io/P2P-Drive/setup"
)

// View Documentation @ http://0.0.0.0:8083/swagger/index.html#/

func main() {
	log.Print("Starting...")
	api := fiber.New()
	setup.LaunchEndpoints(api)
	api.Listen(":8083")
	//test := routes.Connection().GetUserFiles("dom@hotmail.com")
	//fmt.Println(test)
}

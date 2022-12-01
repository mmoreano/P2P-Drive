package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/swagger"
	"log"
	_ "zendx.io/P2P-Drive/docs"
	"zendx.io/P2P-Drive/setup"
)

//  View Documentation @ http://0.0.0.0:8083/swagger/index.html#/

func main() {
	log.Print("Starting...")
	api := fiber.New()
	setup.LaunchEndpoints(api)
	api.Listen(":8083")

	//----Test for Register----
	//testy := models.RegisterRequest{"Dom", "test", "3058332495", "tester@gmail.com", "Dominick", "Diaz", ""}
	//routes.Connection().DBregister(&testy)

	//----Test for login---
	//user := models.LoginRequest{"Dom", "test"}
	//test := routes.Connection().Login(&user)
	//fmt.Println(test)
}

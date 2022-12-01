package main

import (
	_ "github.com/gofiber/swagger"
	"log"
	_ "zendx.io/P2P-Drive/docs"
	"zendx.io/P2P-Drive/models"
	"zendx.io/P2P-Drive/routes"
)

//  View Documentation @ http://0.0.0.0:8083/swagger/index.html#/

func main() {
	log.Print("Starting...")
	//api := fiber.New()
	//setup.LaunchEndpoints(api)
	//api.Listen(":8083")

	testy := models.RegisterRequest{"Dom", "test", "3058332495", "teesteer@gmail.com", "Dominick", "Diaz", ""}
	routes.Connection().DBregister(&testy)

	//user := models.LoginRequest{"Dom", "test"}
	//test := routes.Connection().Login(&user)
	//fmt.Println(test)
}

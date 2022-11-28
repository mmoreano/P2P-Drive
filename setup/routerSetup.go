package setup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	swagger "github.com/gofiber/swagger"
	_ "zendx.io/P2P-Drive/docs"
	"zendx.io/P2P-Drive/routes"
)

// -------------------------- API Endpoints --------------------------\\
func LaunchEndpoints(api *fiber.App) {
	api.Use(recover.New())
	api.Use(cors.New())
	api.Post("/fileAdd", routes.Add)
	api.Get("/get", routes.Get)
	api.Post("/userRegister", routes.UserRegister)
	api.Get("/swagger/*", swagger.HandlerDefault)
}

package setup

import (
	swagger "github.com/gofiber/swagger"
	//"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"zendx.io/P2P-Drive/routes"
	// swagger embed files
)

// -------------------------- API Endpoints --------------------------\\
func LaunchEndpoints(api *fiber.App) {
	api.Use(recover.New())
	api.Use(cors.New())
	api.Post("/add", routes.Add)
	api.Get("/get", routes.Get)
	api.Post("/userRegisters", routes.UserRegister)
	api.Get("/swagger/*", swagger.HandlerDefault)
}

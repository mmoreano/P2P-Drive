package setup

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	swagger "github.com/gofiber/swagger"
	_ "zendx.io/P2P-Drive/docs"
	"zendx.io/P2P-Drive/routes"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8083
// @BasePath /

// -------------------------- API Endpoints --------------------------\\
func LaunchEndpoints(api *fiber.App) {
	api.Use(recover.New())
	api.Use(cors.New())
	api.Post("/add", routes.Add)
	api.Get("/get", routes.Get)
	api.Post("/userRegisters", routes.UserRegister)
	api.Get("/swagger/*", swagger.HandlerDefault)
}

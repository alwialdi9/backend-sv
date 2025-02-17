package routers

import (
	// "github.com/alwialdi9/backend-sv/internal/controllers"
	"github.com/alwialdi9/backend-sv/internal/controllers"
	"github.com/alwialdi9/backend-sv/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/gofiber/swagger"
)

func Route() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(middlewares.LoggerMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Swagger endpoint
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	api := app.Group("/api")

	api.Post("/article", controllers.CreateArticle)
	api.Get("/article/:limit/:offset", controllers.GetArticle)
	api.Get("/article/:id", controllers.GetArticleById)

	return app
}

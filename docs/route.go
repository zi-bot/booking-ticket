package docs

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/docs*", swagger.HandlerDefault)
}

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/qinains/fastergoding"
	"goers/configs"
	"goers/docs"
	_ "goers/docs"
	"goers/pkg/event"
	"goers/pkg/order"
	"goers/pkg/user"
)

func main() {
	fastergoding.Run()
	database := configs.NewDatabase(configs.New())
	validate := validator.New()

	app := fiber.New(configs.NewFiberConfig())
	app.Use(recover.New())

	docs.Router(app)
	user.InitializeUser(database, validate, app)
	event.InitializeEvent(database, validate, app)
	order.InitializeOrder(database, validate, app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}

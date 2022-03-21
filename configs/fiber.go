package configs

import (
	"github.com/gofiber/fiber/v2"
	"goers/execption"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: execption.ErrorHandler,
	}
}

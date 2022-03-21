package execption

import (
	"github.com/gofiber/fiber/v2"
)

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewNotFoundError(err error) {
	if err != nil {
		panic(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}
}

func NewBadRequestError(err error) {
	if err != nil {
		panic(fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}
}

package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"os"
)

type TokenDetails struct {
	UserID   int64   `json:"user_id"`
	UserName string  `json:"user_name"`
	Name     string  `json:"name"`
	Exp      float64 `json:"exp"`
}

func Protected() fiber.Handler {
	config := jwtware.Config{
		SigningKey:   []byte(os.Getenv("SECRET_KEY")),
		ContextKey:   "user",
		ErrorHandler: errorHandler,
	}
	return jwtware.New(config)
}

func errorHandler(c *fiber.Ctx, err error) error {
	err = fiber.NewError(401, err.Error())
	return err
}

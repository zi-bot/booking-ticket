package execption

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"goers/utility"
	"net/http"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	logrus.Error(err)
	_, ok := err.(validator.ValidationErrors)
	if ok {
		return c.Status(http.StatusBadRequest).JSON(utility.WebResponse{
			Code:    400,
			Status:  false,
			Message: http.StatusText(400),
			Data:    nil,
			Error:   err.Error(),
		})
	}

	e, ok := err.(*fiber.Error)
	if ok {
		c.Status(e.Code)
		return c.JSON(utility.WebResponse{
			Code:    e.Code,
			Status:  false,
			Message: http.StatusText(e.Code),
			Data:    nil,
			Error:   err.Error(),
		})
	}
	c.Status(500)
	return c.JSON(utility.WebResponse{
		Code:    500,
		Status:  false,
		Message: http.StatusText(500),
		Data:    nil,
		Error:   err.Error(),
	})
}

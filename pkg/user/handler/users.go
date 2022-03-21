package handler

import (
	"github.com/gofiber/fiber/v2"
	"goers/execption"
	"goers/middleware"
	"goers/pkg/user/model"
	"goers/pkg/user/service"
	"goers/utility"
)

type userHandler struct {
	service  service.UserService
	response *utility.WebResponse
}

func NewUserHandler(UsersService *service.UserService) *userHandler {
	var response utility.WebResponse
	return &userHandler{
		service:  *UsersService,
		response: &response,
	}
}

func (handler *userHandler) Router(app *fiber.App) {
	user := app.Group("/user")
	user.Post("/register", handler.Register)
	user.Post("/login", handler.Login)
	user.Get("/test-protected", middleware.Protected(), handler.TestProtected)
}

func (handler userHandler) TestProtected(c *fiber.Ctx) error {
	ctx := c.Context()
	dataToken, err := utility.ExtractTokenMetadata(ctx)
	execption.PanicError(err)
	return c.JSON(utility.WebResponse{
		Status:  false,
		Code:    0,
		Message: "",
		Data:    dataToken,
		Error:   nil,
	})
}

func (handler *userHandler) Register(c *fiber.Ctx) error {
	var request model.UserRequest
	err := c.BodyParser(&request)
	execption.NewBadRequestError(err)

	err = handler.service.Create(c.Context(), &request)
	execption.NewBadRequestError(err)

	return c.JSON(handler.response.Default(nil, 201))
}

func (handler *userHandler) Login(c *fiber.Ctx) error {
	var request model.UserLoginRequest
	err := c.BodyParser(&request)
	execption.NewBadRequestError(err)

	response, err := handler.service.Login(c.Context(), &request)
	execption.NewBadRequestError(err)

	return c.JSON(handler.response.Default(response, 200))
}

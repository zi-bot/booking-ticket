package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"goers/execption"
	"goers/middleware"
	"goers/pkg/order/model"
	"goers/pkg/order/service"
	"goers/utility"
	"gorm.io/gorm"
)

type orderHandler struct {
	service  service.OrderService
	response *utility.WebResponse
}

func NewOrderHandler(service service.OrderService) *orderHandler {
	var response utility.WebResponse
	return &orderHandler{
		service:  service,
		response: &response,
	}
}

func (h *orderHandler) Router(app *fiber.App) {
	order := app.Group("/order", middleware.Protected())
	{
		order.Get("/", h.List)
		order.Get("/:id", h.Detail)
		order.Post("/", h.Create)
		order.Put("/:id", h.Update)
		order.Delete("/:id", h.Delete)
	}
}

func (h *orderHandler) List(c *fiber.Ctx) error {
	result, err := h.service.GetOrders(c.Context())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)
	return c.JSON(h.response.Default(result, 200))
}

func (h *orderHandler) Detail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	execption.NewBadRequestError(err)

	result, err := h.service.GetOrderByID(c.Context(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)
	return c.JSON(h.response.Default(result, 200))
}

func (h *orderHandler) Create(c *fiber.Ctx) error {
	var request *model.OrderRequest
	err := c.BodyParser(&request)
	execption.NewBadRequestError(err)

	result, err := h.service.CreateOrder(c.Context(), request)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)
	return c.Status(201).JSON(h.response.Default(result, 201))
}

func (h *orderHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	execption.NewBadRequestError(err)

	var request *model.OrderUpdateRequest
	err = c.BodyParser(&request)
	execption.NewBadRequestError(err)

	err = h.service.UpdateOrder(c.Context(), id, request)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)
	return c.JSON(h.response.Default(nil, 200))
}

func (h *orderHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	execption.NewBadRequestError(err)

	err = h.service.DeleteOrder(c.Context(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)
	return c.Status(204).JSON(h.response.Default(nil, 204))
}

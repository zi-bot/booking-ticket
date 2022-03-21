package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"goers/execption"
	"goers/middleware"
	"goers/pkg/event/model"
	"goers/pkg/event/service"
	"goers/utility"
	"gorm.io/gorm"
)

type eventHandler struct {
	service  service.EventService
	response utility.WebResponse
}

func NewEventHandler(service service.EventService) *eventHandler {
	var response utility.WebResponse
	return &eventHandler{
		service:  service,
		response: response,
	}
}

func (h *eventHandler) Router(app *fiber.App) {
	event := app.Group("/event")
	{
		event.Get("/", h.List)
		event.Get("/:id/detail", h.Detail)
		event.Post("/", middleware.Protected(), h.Create)
		event.Put("/:id", middleware.Protected(), h.Update)
		event.Delete("/:id", middleware.Protected(), h.Delete)
		event.Get("/me", middleware.Protected(), h.MyEvent)
	}
}

func (h *eventHandler) List(c *fiber.Ctx) error {
	result, err := h.service.GetEvents(c.Context())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)

	return c.JSON(h.response.Default(result, 200))
}

func (h *eventHandler) Detail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	execption.PanicError(err)

	result, err := h.service.GetEvent(c.Context(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(errors.New("event not found"))
	}
	execption.PanicError(err)

	return c.JSON(h.response.Default(result, 200))
}

func (h *eventHandler) Create(c *fiber.Ctx) error {
	var request model.EventRequest
	err := c.BodyParser(&request)
	execption.NewBadRequestError(err)

	err = h.service.CreateEvent(c.Context(), &request)
	execption.PanicError(err)

	return c.Status(201).JSON(h.response.Default(nil, 201))
}

func (h *eventHandler) Update(c *fiber.Ctx) error {
	var request model.EventUpdateRequest
	id, err := c.ParamsInt("id")
	execption.NewBadRequestError(err)

	err = c.BodyParser(&request)
	execption.NewBadRequestError(err)

	err = h.service.UpdateEvent(c.Context(), id, &request)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(errors.New("event not found"))
	}
	execption.PanicError(err)

	return c.JSON(h.response.Default(nil, 200))
}

func (h *eventHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	execption.NewBadRequestError(err)

	err = h.service.DeleteEvent(c.Context(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(errors.New("event not found"))
	}
	execption.PanicError(err)

	return c.JSON(h.response.Default(nil, 204))
}

func (h *eventHandler) MyEvent(c *fiber.Ctx) error {
	result, err := h.service.GetMyEvents(c.Context())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)
	return c.JSON(h.response.Default(result, 200))
}

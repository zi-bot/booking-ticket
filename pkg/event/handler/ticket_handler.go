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
	"strconv"
)

type ticketHandler struct {
	service  service.TicketService
	response *utility.WebResponse
}

func NewTicketHandler(service service.TicketService) *ticketHandler {
	var response utility.WebResponse
	return &ticketHandler{
		service:  service,
		response: &response,
	}
}

func (h *ticketHandler) Router(app *fiber.App) {
	ticket := app.Group("/ticket")
	{
		ticket.Get("/", h.List)
		ticket.Get("/:id", h.Detail)
		ticket.Post("/", middleware.Protected(), h.Create)
		ticket.Put("/:id", middleware.Protected(), h.Update)
		ticket.Delete("/:id", middleware.Protected(), h.Delete)
	}

}

func (h *ticketHandler) List(c *fiber.Ctx) error {
	param := c.Query("event_id")
	if param == "" {
		param = "0"
	}
	eventID, err := strconv.Atoi(param)
	if err != nil {
		execption.NewBadRequestError(errors.New("event_id is not valid"))
	}

	result, err := h.service.GetTickets(c.Context(), eventID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)

	return c.JSON(h.response.Default(result, 200))
}

func (h *ticketHandler) Detail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	execption.NewBadRequestError(err)

	result, err := h.service.GetTicket(c.Context(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)

	return c.JSON(h.response.Default(result, 200))
}

func (h *ticketHandler) Create(c *fiber.Ctx) error {
	var ticket model.TicketRequest
	err := c.BodyParser(&ticket)
	execption.PanicError(err)

	err = h.service.CreateTicket(c.Context(), &ticket)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)

	return c.JSON(h.response.Default(nil, 201))
}

func (h *ticketHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	execption.NewBadRequestError(err)

	var ticket model.TicketUpdateRequest
	err = c.BodyParser(&ticket)
	execption.PanicError(err)

	err = h.service.UpdateTicket(c.Context(), id, &ticket)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)

	return c.JSON(h.response.Default(nil, 200))
}

func (h *ticketHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	execption.NewBadRequestError(err)

	err = h.service.DeleteTicket(c.Context(), id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		execption.NewNotFoundError(err)
	}
	execption.PanicError(err)

	return c.JSON(h.response.Default(nil, 204))
}

package event

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"goers/pkg/event/handler"
	"goers/pkg/event/repository"
	"goers/pkg/event/service"
	"gorm.io/gorm"
)

func InitializeEvent(database *gorm.DB, validate *validator.Validate, app *fiber.App) {
	eventRepo := repository.NewEventRepositoryImpl(database)
	eventService := service.NewEventServiceImpl(eventRepo, validate)
	eventHandler := handler.NewEventHandler(eventService)
	eventHandler.Router(app)

	ticketRepo := repository.NewTicketRepositoryImpl(database)
	ticketService := service.NewTicketServiceImpl(ticketRepo, eventRepo, validate)
	ticketHandler := handler.NewTicketHandler(ticketService)
	ticketHandler.Router(app)
}

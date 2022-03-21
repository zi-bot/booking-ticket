package order

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	repository2 "goers/pkg/event/repository"
	"goers/pkg/order/handler"
	"goers/pkg/order/repository"
	"goers/pkg/order/service"
	"gorm.io/gorm"
)

func InitializeOrder(database *gorm.DB, validate *validator.Validate, app *fiber.App) {
	orderRepo := repository.NewOrderRepositoryImpl(database)
	ticketRepo := repository2.NewTicketRepositoryImpl(database)
	orderService := service.NewOrderServiceImpl(orderRepo, validate, ticketRepo)
	orderHandler := handler.NewOrderHandler(orderService)
	orderHandler.Router(app)
}

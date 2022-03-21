package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"goers/pkg/user/handler"
	"goers/pkg/user/repository"
	"goers/pkg/user/service"
	"gorm.io/gorm"
)

func InitializeUser(database *gorm.DB, validate *validator.Validate, app *fiber.App) {
	userRepo := repository.NewUsersRepository(database)
	userService := service.NewUserService(&userRepo, validate)
	userHandler := handler.NewUserHandler(&userService)
	userHandler.Router(app)
}

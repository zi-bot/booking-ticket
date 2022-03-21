package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"goers/pkg/user/model"
	"goers/pkg/user/model/entity"
	"goers/pkg/user/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(user *entity.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_name"] = user.UserName
	claims["name"] = user.Name
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

type UserService interface {
	Create(ctx context.Context, request *model.UserRequest) error
	Login(ctx context.Context, request *model.UserLoginRequest) (*model.LoginResponse, error)
	Update(ctx context.Context, id int, request *model.UserRequest) error
}

type userService struct {
	repository repository.UsersRepository
	validate   *validator.Validate
}

func NewUserService(usersRepository *repository.UsersRepository, validate *validator.Validate) UserService {
	return &userService{
		repository: *usersRepository,
		validate:   validate,
	}
}

func (service *userService) Create(ctx context.Context, request *model.UserRequest) error {
	err := service.validate.Struct(request)
	if err != nil {
		return err
	}
	password, err := hashPassword(request.Password)
	if err != nil {
		return err
	}

	_, err = service.repository.GetByUserName(ctx, request.UserName)
	if err == nil {
		return errors.New("username is already taken")
	}
	request.Password = password
	err = service.repository.Create(ctx, request.ToEntity())
	return err
}

func (service *userService) Login(ctx context.Context, request *model.UserLoginRequest) (response *model.LoginResponse, err error) {
	err = service.validate.Struct(request)
	if err != nil {
		return nil, err
	}
	result, err := service.repository.GetByUserName(ctx, request.UserName)
	if err != nil {
		return nil, errors.New("username or password is wrong")
	}
	if !checkPassword(request.Password, result.Password) {
		return nil, errors.New("password is incorrect")
	}
	token, err := generateJWT(result)
	if err != nil {
		return nil, err
	}

	return response.Response(result, token), nil
}

func (service *userService) Update(ctx context.Context, id int, request *model.UserRequest) error {
	//TODO implement me
	panic("implement me")
}

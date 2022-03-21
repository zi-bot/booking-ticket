package model

import "goers/pkg/user/model/entity"

type UserRequest struct {
	Name     string `validate:"required,max=45,min=1" json:"name"`
	UserName string `validate:"required,max=45,min=1" json:"user_name"`
	Password string `validate:"required" json:"password"`
}

func (r *UserRequest) ToEntity() *entity.User {
	return &entity.User{
		Name:     r.Name,
		UserName: r.UserName,
		Password: r.Password,
	}
}

type UserLoginRequest struct {
	UserName string `validate:"required,max=45,min=1" json:"user_name"`
	Password string `validate:"required" json:"password"`
}

func (r *UserLoginRequest) ToEntity() *entity.User {
	return &entity.User{
		UserName: r.UserName,
		Password: r.Password,
	}
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
}

func (r *UserResponse) Response(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		UserName: user.UserName,
	}
}

type LoginResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"access_token"`
}

func (r *LoginResponse) Response(user *entity.User, token string) *LoginResponse {
	var userResponse UserResponse
	return &LoginResponse{
		User:        *userResponse.Response(user),
		AccessToken: token,
	}
}

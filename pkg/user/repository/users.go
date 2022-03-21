package repository

import (
	"context"
	"goers/pkg/user/model/entity"
	"gorm.io/gorm"
)

type UsersRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetById(ctx context.Context, id int) (*entity.User, error)
	GetByUserName(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, id int, user *entity.User) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) (*[]entity.User, error)
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(database *gorm.DB) UsersRepository {
	return &usersRepository{
		db: database,
	}
}

func (r *usersRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *usersRepository) GetById(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *usersRepository) GetByUserName(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("user_name = ?", username).First(&user).Error
	return &user, err
}

func (r *usersRepository) Update(ctx context.Context, id int, user *entity.User) error {
	var userEntity entity.User
	res := r.db.First(&userEntity, id).UpdateColumns(user)
	return res.Error
}

func (r *usersRepository) Delete(ctx context.Context, id int) error {
	var user entity.User
	res := r.db.First(&user, id).Delete(&user)
	return res.Error
}

func (r *usersRepository) GetAll(ctx context.Context) (*[]entity.User, error) {
	var users []entity.User
	res := r.db.Find(&users)
	if len(users) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &users, res.Error
}

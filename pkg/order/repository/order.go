package repository

import (
	"context"
	"goers/pkg/order/model/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAll(ctx context.Context) (*[]entity.Order, error)
	GetByID(ctx context.Context, id int) (*entity.Order, error)
	GetByOwner(ctx context.Context, ownerID int) (*[]entity.Order, error)
	GetByOwnerAndID(ctx context.Context, ownerID int, id int) (*entity.Order, error)
	Create(ctx context.Context, order *entity.Order) (*entity.Order, error)
	Update(ctx context.Context, id int, order *entity.Order) error
	Delete(ctx context.Context, id int) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepositoryImpl(database *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{
		db: database,
	}
}

func (r *orderRepositoryImpl) GetAll(ctx context.Context) (*[]entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepositoryImpl) GetByID(ctx context.Context, id int) (order *entity.Order, err error) {
	res := r.db.WithContext(ctx).Preload("Ticket").Preload("Owner").First(&order, id)
	return order, res.Error
}

func (r *orderRepositoryImpl) GetByOwner(ctx context.Context, ownerID int) (orders *[]entity.Order, err error) {
	res := r.db.WithContext(ctx).Preload("Owner").Preload("Ticket").Where("owner_id = ?", ownerID).Find(&orders)
	return orders, res.Error
}

func (r *orderRepositoryImpl) Create(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	res := r.db.WithContext(ctx).Create(order)
	return order, res.Error
}

func (r *orderRepositoryImpl) Update(ctx context.Context, id int, order *entity.Order) error {
	var entityOrder *entity.Order
	res := r.db.WithContext(ctx).First(entityOrder, id).UpdateColumns(order)
	return res.Error
}

func (r *orderRepositoryImpl) Delete(ctx context.Context, id int) error {
	var entityOrder *entity.Order
	res := r.db.WithContext(ctx).First(entityOrder, id).Delete(&entityOrder)
	return res.Error
}

func (r *orderRepositoryImpl) GetByOwnerAndID(ctx context.Context, ownerID int, id int) (order *entity.Order, err error) {
	res := r.db.WithContext(ctx).Preload("Owner").Preload("Ticket").Where("owner_id = ? AND id = ?", ownerID, id).First(&order)
	return order, res.Error
}

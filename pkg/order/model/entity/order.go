package entity

import (
	entity2 "goers/pkg/event/model/entity"
	"goers/pkg/user/model/entity"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusPending  OrderStatus = "pending"
	OrderStatusCanceled OrderStatus = "canceled"
	OrderStatusPaid     OrderStatus = "paid"
	OrderStatusDone     OrderStatus = "done"
)

type Order struct {
	gorm.Model
	OrderNumber string         `gorm:"type:varchar(255);unique_index"`
	TicketID    uint           `gorm:"not null"`
	Ticket      entity2.Ticket `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OwnerID     uint           `gorm:"not null;index"`
	Owner       entity.User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TotalPrice  float64        `gorm:"not null"`
	Qty         int            `gorm:"not null"`
	Status      OrderStatus    `gorm:"not null;index"`
}

package model

import (
	"goers/pkg/event/model"
	"goers/pkg/order/model/entity"
	model2 "goers/pkg/user/model"
	"time"
)

type OrderResponse struct {
	ID          uint                  `json:"id"`
	OrderNumber string                `json:"order_number"`
	TicketID    uint                  `json:"ticket_id"`
	Ticket      *model.TicketResponse `json:"ticket"`
	OwnerID     uint                  `json:"owner_id"`
	Owner       *model2.UserResponse  `json:"owner"`
	TotalPrice  float64               `json:"total_price"`
	Qty         int                   `json:"qty"`
	Status      entity.OrderStatus    `json:"status"`
}

func (r *OrderResponse) Default(order *entity.Order) *OrderResponse {

	return &OrderResponse{
		ID:          order.ID,
		OrderNumber: order.OrderNumber,
		TicketID:    order.TicketID,
		OwnerID:     order.OwnerID,
		TotalPrice:  order.TotalPrice,
		Qty:         order.Qty,
		Status:      order.Status,
	}
}

func (r *OrderResponse) ResponseList(entityOrders *[]entity.Order) *[]OrderResponse {
	var orders []OrderResponse
	for _, order := range *entityOrders {
		orders = append(orders, *r.Default(&order))
	}
	return &orders
}

func (r *OrderResponse) ResponseDetail(entityOrder *entity.Order) *OrderResponse {
	var ticket model.TicketResponse
	var owner model2.UserResponse
	ownerResponse := owner.Response(&entityOrder.Owner)
	ticketResponse := ticket.Default(&entityOrder.Ticket)
	return &OrderResponse{
		ID:          entityOrder.ID,
		OrderNumber: entityOrder.OrderNumber,
		TicketID:    entityOrder.TicketID,
		Ticket:      ticketResponse,
		OwnerID:     entityOrder.OwnerID,
		Owner:       ownerResponse,
		TotalPrice:  entityOrder.TotalPrice,
		Qty:         entityOrder.Qty,
		Status:      entityOrder.Status,
	}

}

type OrderRequest struct {
	TicketID uint `validate:"required" json:"ticket_id"`
	Qty      int  `validate:"required,min=1" json:"qty"`
}

func (r *OrderRequest) ToEntity(totalPrice float64, ownerID uint) *entity.Order {
	var orderNumber string

	orderNumber = "INV-" + time.Now().Format("20060102150405")

	return &entity.Order{
		OrderNumber: orderNumber,
		OwnerID:     ownerID,
		TicketID:    r.TicketID,
		TotalPrice:  totalPrice,
		Qty:         r.Qty,
		Status:      entity.OrderStatusPending,
	}
}

type OrderUpdateRequest struct {
	TicketID uint               `validate:"omitempty,min=1" json:"ticket_id"`
	Qty      int                `validate:"omitempty,min=1" json:"qty"`
	Status   entity.OrderStatus `validate:"omitempty,oneof='pending' 'canceled' 'paid' 'done'"json:"status"`
}

func (r OrderUpdateRequest) ToEntity() *entity.Order {
	return &entity.Order{
		TicketID: r.TicketID,
		Qty:      r.Qty,
		Status:   r.Status,
	}
}

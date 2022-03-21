package model

import (
	"goers/pkg/event/model/entity"
)

type TicketRequest struct {
	EventID uint    `validate:"required" json:"event_id"`
	Price   float64 `validate:"required" json:"price"`
	Detail  string  `validate:"required" json:"detail"`
	Name    string  `validate:"required" json:"name"`
	Stock   int     `validate:"required" json:"stock"`
}

func (r *TicketRequest) ToEntity() *entity.Ticket {
	return &entity.Ticket{
		EventID: r.EventID,
		Price:   r.Price,
		Detail:  r.Detail,
		Name:    r.Name,
		Stock:   r.Stock,
	}
}

type TicketUpdateRequest struct {
	EventID uint    `validate:"omitempty" json:"event_id"`
	Price   float64 `validate:"omitempty" json:"price"`
	Detail  string  `validate:"omitempty" json:"detail"`
	Name    string  `validate:"omitempty" json:"name"`
	Stock   int     `validate:"omitempty" json:"quantity"`
}

func (r *TicketUpdateRequest) ToEntity() *entity.Ticket {
	return &entity.Ticket{
		EventID: r.EventID,
		Price:   r.Price,
		Detail:  r.Detail,
		Name:    r.Name,
		Stock:   r.Stock,
	}
}

type TicketResponse struct {
	ID      uint           `json:"id"`
	EventID uint           `json:"event_id"`
	Price   float64        `json:"price"`
	Detail  string         `json:"detail"`
	Name    string         `json:"name"`
	Stock   int            `json:"quantity"`
	Event   *EventResponse `json:"event"`
}

func (r *TicketResponse) Default(entityTicket *entity.Ticket) *TicketResponse {
	return &TicketResponse{
		ID:      entityTicket.ID,
		EventID: entityTicket.EventID,
		Price:   entityTicket.Price,
		Detail:  entityTicket.Detail,
		Name:    entityTicket.Name,
		Stock:   entityTicket.Stock,
	}
}

func (r *TicketResponse) ResponseList(entityTickets *[]entity.Ticket) *[]TicketResponse {
	var tickets []TicketResponse
	for _, entityTicket := range *entityTickets {
		tickets = append(tickets, *r.Default(&entityTicket))
	}
	return &tickets
}

func (r *TicketResponse) ResponseDetail(entityTicket *entity.Ticket) *TicketResponse {
	var eventResponse *EventResponse
	return &TicketResponse{
		ID:      entityTicket.ID,
		EventID: entityTicket.EventID,
		Price:   entityTicket.Price,
		Detail:  entityTicket.Detail,
		Name:    entityTicket.Name,
		Stock:   entityTicket.Stock,
		Event:   eventResponse.Default(&entityTicket.Event),
	}
}

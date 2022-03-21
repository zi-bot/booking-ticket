package repository

import (
	"context"
	"goers/pkg/event/model/entity"
	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ctx context.Context, ticket *entity.Ticket) error
	Update(ctx context.Context, id int, ticket *entity.Ticket) error
	Delete(ctx context.Context, id int) error
	GetTicketsByEventID(ctx context.Context, eventID int) (*[]entity.Ticket, error)
	GetByID(ctx context.Context, id int) (*entity.Ticket, error)
	GetAll(ctx context.Context) (*[]entity.Ticket, error)
}

type ticketImpl struct {
	db *gorm.DB
}

func NewTicketRepositoryImpl(db *gorm.DB) TicketRepository {
	return &ticketImpl{
		db: db,
	}
}

func (r *ticketImpl) Create(ctx context.Context, ticket *entity.Ticket) error {
	res := r.db.WithContext(ctx).Create(ticket)
	return res.Error
}

func (r *ticketImpl) Update(ctx context.Context, id int, ticket *entity.Ticket) error {
	var entityTicket = entity.Ticket{}
	res := r.db.WithContext(ctx).First(&entityTicket, id).UpdateColumns(ticket)
	return res.Error
}

func (r *ticketImpl) Delete(ctx context.Context, id int) error {
	var ticket entity.Ticket
	res := r.db.WithContext(ctx).First(&ticket, id).Delete(&ticket)
	return res.Error
}

func (r *ticketImpl) GetTicketsByEventID(ctx context.Context, eventID int) (ticket *[]entity.Ticket, err error) {
	res := r.db.WithContext(ctx).Preload("Event").Where("event_id = ?", eventID).Find(&ticket)
	return ticket, res.Error
}

func (r *ticketImpl) GetAll(ctx context.Context) (tickets *[]entity.Ticket, err error) {
	res := r.db.WithContext(ctx).Preload("Event").Find(&tickets)
	return tickets, res.Error
}

func (r *ticketImpl) GetByID(ctx context.Context, id int) (ticket *entity.Ticket, err error) {
	res := r.db.WithContext(ctx).Preload("Event").First(&ticket, id)
	return ticket, res.Error
}

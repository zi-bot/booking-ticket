package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	repository2 "goers/pkg/event/repository"
	"goers/pkg/order/model"
	"goers/pkg/order/repository"
	"goers/utility"
)

type OrderService interface {
	GetOrderByID(ctx context.Context, id int) (*model.OrderResponse, error)
	GetOrders(ctx context.Context) (*[]model.OrderResponse, error)
	CreateOrder(ctx context.Context, order *model.OrderRequest) (*model.OrderResponse, error)
	UpdateOrder(ctx context.Context, id int, order *model.OrderUpdateRequest) error
	DeleteOrder(ctx context.Context, id int) error
}

type orderServiceImpl struct {
	repository       repository.OrderRepository
	ticketRepository repository2.TicketRepository
	validate         *validator.Validate
}

func NewOrderServiceImpl(repository repository.OrderRepository, validate *validator.Validate,
	ticketRepository repository2.TicketRepository) OrderService {
	return &orderServiceImpl{
		repository:       repository,
		ticketRepository: ticketRepository,
		validate:         validate,
	}
}

func (s *orderServiceImpl) GetOrderByID(ctx context.Context, id int) (response *model.OrderResponse, err error) {
	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.repository.GetByOwnerAndID(ctx, int(token.UserID), id)
	return response.ResponseDetail(result), err
}

func (s *orderServiceImpl) GetOrders(ctx context.Context) (*[]model.OrderResponse, error) {
	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return nil, err
	}

	result, err := s.repository.GetByOwner(ctx, int(token.UserID))
	var response *model.OrderResponse
	return response.ResponseList(result), err
}

func (s *orderServiceImpl) CreateOrder(ctx context.Context, order *model.OrderRequest) (response *model.OrderResponse, err error) {
	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return nil, err
	}
	//	check event and ticket
	err = s.validate.Struct(order)
	if err != nil {
		return nil, err
	}
	ticket, err := s.ticketRepository.GetByID(ctx, int(order.TicketID))
	if err != nil {
		return nil, err
	}
	if order.Qty >= ticket.Stock {
		return nil, errors.New("stock is not enough")
	}

	ticket.Stock = ticket.Stock - order.Qty
	err = s.ticketRepository.Update(ctx, int(ticket.ID), ticket)
	if err != nil {
		return nil, err
	}
	var totalPrice float64
	totalPrice = ticket.Price * float64(order.Qty)
	result, err := s.repository.Create(ctx, order.ToEntity(totalPrice, uint(token.UserID)))
	return response.Default(result), err
}

func (s *orderServiceImpl) UpdateOrder(ctx context.Context, id int, order *model.OrderUpdateRequest) (err error) {
	_, err = s.repository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = s.validate.Struct(order)
	if err != nil {
		return err
	}

	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return err
	}
	ticket, err := s.ticketRepository.GetByID(ctx, int(order.TicketID))
	if ticket.Event.OwnerID != uint(token.UserID) {
		return errors.New("you are not owner of this event")
	}
	err = s.repository.Update(ctx, id, order.ToEntity())
	return
}

func (s *orderServiceImpl) DeleteOrder(ctx context.Context, id int) error {
	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return err
	}
	_, err = s.repository.GetByOwnerAndID(ctx, int(token.UserID), id)
	return s.repository.Delete(ctx, id)
}

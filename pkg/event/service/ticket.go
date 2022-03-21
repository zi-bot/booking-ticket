package service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"goers/pkg/event/model"
	"goers/pkg/event/repository"
	"goers/utility"
)

type TicketService interface {
	GetTicket(ctx context.Context, id int) (*model.TicketResponse, error)
	GetTickets(ctx context.Context, eventID int) (*[]model.TicketResponse, error)
	CreateTicket(ctx context.Context, ticket *model.TicketRequest) error
	UpdateTicket(ctx context.Context, id int, ticket *model.TicketUpdateRequest) error
	DeleteTicket(ctx context.Context, id int) error
}

type ticketServiceImpl struct {
	ticketRepository repository.TicketRepository
	validate         *validator.Validate
	eventRepository  repository.EventRepository
}

func NewTicketServiceImpl(ticketRepository repository.TicketRepository, eventRepository repository.EventRepository,
	validate *validator.Validate) TicketService {
	return &ticketServiceImpl{
		ticketRepository: ticketRepository,
		eventRepository:  eventRepository,
		validate:         validate,
	}
}

func (s *ticketServiceImpl) GetTicket(ctx context.Context, id int) (response *model.TicketResponse, err error) {
	result, err := s.ticketRepository.GetByID(ctx, id)
	return response.ResponseDetail(result), err
}

func (s *ticketServiceImpl) GetTickets(ctx context.Context, eventID int) (*[]model.TicketResponse, error) {

	var response *model.TicketResponse

	if eventID != 0 {
		event, err := s.eventRepository.GetByID(ctx, eventID)
		if err != nil {
			return nil, err
		}
		result, err := s.ticketRepository.GetTicketsByEventID(ctx, int(event.ID))
		if err != nil {
			return nil, err
		}
		return response.ResponseList(result), err

	}
	result, err := s.ticketRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return response.ResponseList(result), err
}

func (s *ticketServiceImpl) CreateTicket(ctx context.Context, ticket *model.TicketRequest) error {
	err := s.validate.Struct(ticket)
	if err != nil {
		return err
	}
	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return err
	}
	_, err = s.eventRepository.GetByIDAndOwner(ctx, int(ticket.EventID), int(token.UserID))
	if err != nil {
		return err
	}
	err = s.ticketRepository.Create(ctx, ticket.ToEntity())
	return err
}

func (s *ticketServiceImpl) UpdateTicket(ctx context.Context, id int, ticket *model.TicketUpdateRequest) error {
	err := s.validate.Struct(ticket)
	if err != nil {
		return err
	}
	fmt.Println(ticket.EventID)
	_, err = s.ticketRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if ticket.EventID > 0 {
		token, err := utility.ExtractTokenMetadata(ctx)
		if err != nil {
			return err
		}
		_, err = s.eventRepository.GetByIDAndOwner(ctx, int(ticket.EventID), int(token.UserID))
		if err != nil {
			return err
		}
	}

	err = s.ticketRepository.Update(ctx, id, ticket.ToEntity())
	return err
}

func (s *ticketServiceImpl) DeleteTicket(ctx context.Context, id int) error {
	ticket, err := s.ticketRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return err
	}
	_, err = s.eventRepository.GetByIDAndOwner(ctx, int(ticket.EventID), int(token.UserID))
	if err != nil {
		return err
	}
	err = s.ticketRepository.Delete(ctx, id)
	return err
}

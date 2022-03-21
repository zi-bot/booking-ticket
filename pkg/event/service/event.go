package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"goers/pkg/event/model"
	"goers/pkg/event/repository"
	"goers/utility"
)

type EventService interface {
	GetEvent(ctx context.Context, id int) (*model.EventResponse, error)
	GetEvents(ctx context.Context) (*[]model.EventResponse, error)
	GetMyEvents(ctx context.Context) (*[]model.EventResponse, error)
	CreateEvent(ctx context.Context, request *model.EventRequest) error
	UpdateEvent(ctx context.Context, id int, request *model.EventUpdateRequest) error
	DeleteEvent(ctx context.Context, id int) error
}

type eventServiceImpl struct {
	repository repository.EventRepository
	validate   *validator.Validate
}

func NewEventServiceImpl(repository repository.EventRepository, validate *validator.Validate) EventService {
	return &eventServiceImpl{
		repository: repository,
		validate:   validate,
	}
}

func (s *eventServiceImpl) GetEvent(ctx context.Context, id int) (response *model.EventResponse, err error) {
	res, err := s.repository.GetByID(ctx, id)
	return response.Default(res), err
}

func (s *eventServiceImpl) GetEvents(ctx context.Context) (responses *[]model.EventResponse, err error) {
	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return
	}
	var response model.EventResponse
	return response.ResponseList(result), err
}

func (s *eventServiceImpl) CreateEvent(ctx context.Context, request *model.EventRequest) (err error) {
	err = s.validate.Struct(request)
	if err != nil {
		return
	}
	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return
	}

	err = s.repository.Create(ctx, request.ToEntity(uint(token.UserID)))
	return
}

func (s *eventServiceImpl) UpdateEvent(ctx context.Context, id int, request *model.EventUpdateRequest) (err error) {
	err = s.validate.Struct(request)
	if err != nil {
		return
	}
	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return
	}
	_, err = s.repository.GetByIDAndOwner(ctx, id, int(token.UserID))
	if err != nil {
		return
	}
	err = s.repository.Update(ctx, id, request.ToEntity())
	return
}

func (s *eventServiceImpl) DeleteEvent(ctx context.Context, id int) (err error) {
	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return
	}
	_, err = s.repository.GetByIDAndOwner(ctx, id, int(token.UserID))
	if err != nil {
		return
	}
	return s.repository.Delete(ctx, id)
}

func (s *eventServiceImpl) GetMyEvents(ctx context.Context) (*[]model.EventResponse, error) {
	token, err := utility.ExtractTokenMetadata(ctx)
	if err != nil {
		return nil, err
	}
	result, err := s.repository.GetByOwner(ctx, int(token.UserID))
	if err != nil {
		return nil, err
	}
	var response model.EventResponse
	return response.ResponseList(result), err
}

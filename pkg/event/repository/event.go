package repository

import (
	"context"
	"goers/pkg/event/model/entity"
	"gorm.io/gorm"
)

type EventRepository interface {
	GetAll(ctx context.Context) (*[]entity.Event, error)
	GetByOwner(ctx context.Context, ownerID int) (*[]entity.Event, error)
	GetByID(ctx context.Context, id int) (*entity.Event, error)
	GetByIDAndOwner(ctx context.Context, id int, ownerID int) (*entity.Event, error)
	Create(ctx context.Context, event *entity.Event) error
	Update(ctx context.Context, id int, event *entity.Event) error
	Delete(ctx context.Context, id int) error
}

type eventImpl struct {
	db *gorm.DB
}

func NewEventRepositoryImpl(db *gorm.DB) EventRepository {
	return &eventImpl{
		db: db,
	}
}

func (r *eventImpl) GetAll(ctx context.Context) (events *[]entity.Event, err error) {
	res := r.db.WithContext(ctx).Find(&events)
	if len(*events) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return events, res.Error
}

func (r *eventImpl) GetByID(ctx context.Context, id int) (event *entity.Event, err error) {
	res := r.db.WithContext(ctx).First(&event, id)
	return event, res.Error
}

func (r *eventImpl) Create(ctx context.Context, event *entity.Event) error {
	res := r.db.WithContext(ctx).Create(event)
	return res.Error
}

func (r *eventImpl) Update(ctx context.Context, id int, event *entity.Event) error {
	var entityEvent = entity.Event{}
	res := r.db.WithContext(ctx).First(&entityEvent, id).UpdateColumns(event)
	return res.Error
}

func (r *eventImpl) Delete(ctx context.Context, id int) error {
	var entityEvent = entity.Event{}
	res := r.db.WithContext(ctx).First(&entityEvent, id).Delete(&entityEvent)
	return res.Error
}

func (r *eventImpl) GetByOwner(ctx context.Context, ownerID int) (events *[]entity.Event, err error) {
	res := r.db.WithContext(ctx).Find(&events, "owner_id = ?", ownerID)
	return events, res.Error
}

func (r *eventImpl) GetByIDAndOwner(ctx context.Context, id int, ownerID int) (event *entity.Event, err error) {
	res := r.db.WithContext(ctx).First(&event, "id = ? AND owner_id = ?", id, ownerID)
	return event, res.Error
}

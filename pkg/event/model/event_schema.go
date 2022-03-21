package model

import (
	"goers/pkg/event/model/entity"
	"goers/utility"
	"time"
)

type EventRequest struct {
	Title       string       `validate:"required" json:"title"`
	Category    string       `validate:"required" json:"category"`
	Description string       `validate:"required" json:"description"`
	Location    string       `validate:"required" json:"location"`
	Latitude    float64      `validate:"required,latitude" json:"latitude"`
	Longitude   float64      `validate:"required,longitude" json:"longitude"`
	StartDate   utility.Date `validate:"required,datetime=2006-01-02" json:"start_date"`
	EndDate     utility.Date `validate:"required,datetime=2006-01-02" json:"end_date"`
}

func (r *EventRequest) ToEntity(ownerID uint) *entity.Event {

	return &entity.Event{
		Title:       r.Title,
		Category:    r.Category,
		Description: r.Description,
		Location:    r.Location,
		Latitude:    r.Latitude,
		Longitude:   r.Longitude,
		StartDate:   r.StartDate.ToTime(),
		EndDate:     r.EndDate.ToTime(),
		OwnerID:     ownerID,
	}
}

type EventUpdateRequest struct {
	Title       string       `validate:"omitempty" json:"title"`
	Category    string       `validate:"omitempty" json:"category"`
	Description string       `validate:"omitempty" json:"description"`
	Location    string       `validate:"omitempty" json:"location"`
	Latitude    float64      `validate:"omitempty,required_with=Longitude,latitude" json:"latitude"`
	Longitude   float64      `validate:"omitempty,required_with=Latitude,longitude" json:"longitude"`
	StartDate   utility.Date `validate:"omitempty,datetime=2006-01-02" json:"start_date"`
	EndDate     utility.Date `validate:"omitempty,datetime=2006-01-02" json:"end_date"`
}

func (r *EventUpdateRequest) ToEntity() *entity.Event {
	return &entity.Event{
		Title:       r.Title,
		Category:    r.Category,
		Description: r.Description,
		Location:    r.Location,
		Latitude:    r.Latitude,
		Longitude:   r.Longitude,
		StartDate:   r.StartDate.ToTime(),
		EndDate:     r.EndDate.ToTime(),
	}
}

type EventResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	OwnerID     uint      `json:"owner_id"`
	Location    string    `json:"location"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

func (r *EventResponse) Default(event *entity.Event) *EventResponse {
	return &EventResponse{
		ID:          event.ID,
		Title:       event.Title,
		Category:    event.Category,
		OwnerID:     event.OwnerID,
		Description: event.Description,
		Location:    event.Location,
		Latitude:    event.Latitude,
		Longitude:   event.Longitude,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
	}
}

func (r *EventResponse) ResponseList(entities *[]entity.Event) *[]EventResponse {
	var events []EventResponse
	for _, entityEvent := range *entities {
		events = append(events, *r.Default(&entityEvent))
	}
	return &events
}

package entity

import (
	"goers/pkg/user/model/entity"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Title       string      `gorm:"not null"`
	Category    string      `gorm:"not null"`
	Description string      `gorm:"text not null"`
	Owner       entity.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OwnerID     uint        `gorm:"not null"`
	Location    string      `gorm:"not null"`
	Latitude    float64     `gorm:"not null"`
	Longitude   float64     `gorm:"not null"`
	StartDate   time.Time   `gorm:"not null"`
	EndDate     time.Time   `gorm:"not null"`
}

package entity

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	EventID uint    `gorm:"not null"`
	Event   Event   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Price   float64 `gorm:"not null"`
	Detail  string  `gorm:"text not null"`
	Name    string  `gorm:"not null"`
	Stock   int     `gorm:"not null"`
}

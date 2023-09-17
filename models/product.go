package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	UUID      string `gorm:"not null"`
	Name      string `gorm:"not null" form:"name" valid:"required~Product name is required."`
	ImageURL  string `gorm:"not null;" valid:"required~Image URL is required."`
	AdminID   uint   `gorm:"not null;" valid:"required~Admin ID is required."`
	Variants  []Variant
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

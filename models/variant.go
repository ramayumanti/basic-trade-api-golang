package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	ID          uint   `gorm:"primaryKey"`
	UUID        string `gorm:"not null"`
	VariantName string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	ProductID   uint   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {

	if len(v.VariantName) < 4 {
		err = errors.New("Variant name is too short")
	}

	if v.Quantity < 0 {
		err = errors.New("Variant stock cannot be less than zero")
	}

	return
}

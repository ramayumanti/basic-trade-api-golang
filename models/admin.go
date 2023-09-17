package models

import (
	"basictrade-api/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint       `gorm:"primaryKey"`
	UUID      string     `gorm:"not null"`
	Name      string     `gorm:"not null" form:"name" valid:"required~Admin name is required."`
	Email     string     `gorm:"not null" json:"email" form:"email" valid:"required~Admin email is required."`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required~Password is required."`
	Products  []Product  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(a)

	if errCreate != nil {
		err = errCreate
		return
	}

	a.Password = helpers.HashPass(a.Password)

	err = nil
	return
}

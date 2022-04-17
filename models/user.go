package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Firtsname string    `json:"first_name" gorm:"type:varchar(255)"`
	Lastname  string    `json:"last_name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" validate:"required,email,min=6,max=32" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"password" gorm:"type:varchar"`

	Varifired bool `json:"is_varified" gorm:"type:boolean"`

	// TODO needs to covert to ENUM type later
	UserType     string `json:"user_type" gorm:"type:varchar(100)"`
	Authprovider string `json:"auth_provider" gorm:"type:varchar(60)"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}

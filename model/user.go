package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Email    *string
	Birthday *time.Time
	Hash     string
	Assets   []Asset `gorm:"many2many:users_pins;"`
}

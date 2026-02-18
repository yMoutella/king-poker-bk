package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
}

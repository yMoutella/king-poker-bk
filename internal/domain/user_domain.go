package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
}

func (p *User) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == (uuid.UUID{}) { // Check if the ID is a zero value UUID
		p.ID, err = uuid.NewV7()
	}
	return err
}

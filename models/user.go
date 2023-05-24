package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FullName  string     `gorm:"type:varchar(255);not null"`
	Email     string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string     `gorm:"type:varchar(255);not null"`
	Role      string     `gorm:"type:varchar(50);not null;default:'user'"`
	CreatedAt *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
}

type UserRepository interface {
	Create(user *User) error
	FetchByID(id string) (User, error)
	FetchByEmail(email string) (User, error)
	FetchAll() ([]User, error)
}

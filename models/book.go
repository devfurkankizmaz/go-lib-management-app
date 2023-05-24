package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID        uuid.UUID  `gorm:"type:uuid;not null"`
	Title         string     `gorm:"type:varchar(255);not null"`
	Author        string     `gorm:"type:varchar(255);not null"`
	Description   string     `gorm:"type:text"`
	CoverPhotoUrl string     `gorm:"type:varchar(255)"`
	CreatedAt     *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
	UpdatedAt     *time.Time `gorm:"type:timestamptz;default:current_timestamp"`
}

type BookInput struct {
	Title         string `json:"title" validate:"required"`
	Author        string `json:"author" validate:"required"`
	Description   string `json:"description"`
	CoverPhotoUrl string `json:"cover_photo_url"`
}

type BookResponse struct {
	Title         string     `json:"title"`
	Author        string     `json:"author"`
	Description   string     `json:"description"`
	CoverPhotoUrl string     `json:"cover_photo_url"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type BookRepository interface {
	Create(book *Book) error
	FetchAllByUserID(id string, limit int, offset int) ([]Book, error)
	FetchByID(id string) (Book, error)
	UpdateByID(book *Book, id string) error
	DeleteByID(id string) error
}

type BookService interface {
	Create(book *Book) error
	FetchAllByUserID(id string, limit int, page int) ([]Book, error)
	FetchByID(id string) (Book, error)
	UpdateByID(book *Book, id string) error
	DeleteByID(id string) error
}

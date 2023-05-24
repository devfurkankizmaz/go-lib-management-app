package repository

import (
	"errors"
	"strings"

	"github.com/devfurkankizmaz/go-lib-management-app/models"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) models.BookRepository {
	return &bookRepository{db: db}
}

func (br *bookRepository) Create(book *models.Book) error {
	result := br.db.Create(&book)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return errors.New("user with that email already exists")
	} else if result.Error != nil {
		return errors.New("something bad happened")
	}
	return nil
}

func (br *bookRepository) FetchAllByUserID(id string, limit int, offset int) ([]models.Book, error) {
	var books = []models.Book{}
	result := br.db.Where("user_id = ?", id).Limit(limit).Offset(offset).Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (br *bookRepository) FetchByID(id string) (models.Book, error) {
	var book = models.Book{}
	result := br.db.Where("id = ?", id).First(&book)

	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

func (br *bookRepository) UpdateByID(book *models.Book, id string) error {
	columbs := map[string]interface{}{"title": book.Title, "author": book.Author, "description": book.Description, "cover_photo_url": book.CoverPhotoUrl}
	result := br.db.Model(&book).Where("id = ?", id).Updates(columbs)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (br *bookRepository) DeleteByID(id string) error {
	result := br.db.Delete(&models.Book{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return result.Error
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}

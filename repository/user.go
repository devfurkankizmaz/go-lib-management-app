package repository

import (
	"errors"
	"strings"

	"github.com/devfurkankizmaz/go-lib-management-app/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(user *models.User) error {
	result := ur.db.Create(&user)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return errors.New("user with that email already exists")
	} else if result.Error != nil {
		return errors.New("something bad happened")
	}
	return nil
}

func (ur *userRepository) FetchAll() ([]models.User, error) {
	var users = []models.User{}
	result := ur.db.Find(&users)

	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

func (ur *userRepository) FetchByID(id string) (models.User, error) {
	var user = models.User{}
	result := ur.db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (ur *userRepository) FetchByEmail(email string) (models.User, error) {
	var user = models.User{}
	result := ur.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

package service

import (
	"github.com/devfurkankizmaz/go-lib-management-app/models"
	"github.com/devfurkankizmaz/go-lib-management-app/utils"
)

type refreshService struct {
	userRepository models.UserRepository
}

func NewRefreshService(userRepository models.UserRepository) models.RefreshService {
	return &refreshService{
		userRepository: userRepository,
	}
}
func (rs *refreshService) FetchByID(id string) (models.User, error) {
	user, err := rs.userRepository.FetchByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (rs *refreshService) GenerateAccessToken(user *models.User, key string, ex int) (token string, err error) {
	return utils.GenerateAccessToken(user, key, ex)
}

func (rs *refreshService) GenerateRefreshToken(user *models.User, key string, ex int) (token string, err error) {
	return utils.GenerateRefreshToken(user, key, ex)
}

func (rs *refreshService) ExtractIDFromToken(inputToken string, key string) (string, error) {
	return utils.ExtractIDFromToken(inputToken, key)
}

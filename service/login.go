package service

import (
	"github.com/devfurkankizmaz/go-lib-management-app/models"
	"github.com/devfurkankizmaz/go-lib-management-app/utils"
)

type loginService struct {
	userRepository models.UserRepository
}

func NewLoginService(repo models.UserRepository) models.LoginService {
	return &loginService{userRepository: repo}
}

func (ls *loginService) GenerateAccessToken(user *models.User, key string, ex int) (token string, err error) {
	return utils.GenerateAccessToken(user, key, ex)
}

func (ls *loginService) GenerateRefreshToken(user *models.User, key string, ex int) (token string, err error) {
	return utils.GenerateRefreshToken(user, key, ex)
}

func (ls *loginService) FetchByEmail(email string) (models.User, error) {
	user, err := ls.userRepository.FetchByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

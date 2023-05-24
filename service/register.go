package service

import "github.com/devfurkankizmaz/go-lib-management-app/models"

type registerService struct {
	userRepository models.UserRepository
}

func NewRegisterService(repo models.UserRepository) models.RegisterService {
	return &registerService{userRepository: repo}
}

func (rs *registerService) Create(user *models.User) error {
	err := rs.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (rs *registerService) FetchByEmail(email string) (models.User, error) {
	user, err := rs.userRepository.FetchByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

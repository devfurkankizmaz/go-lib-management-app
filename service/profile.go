package service

import "github.com/devfurkankizmaz/go-lib-management-app/models"

type profileService struct {
	userRepository models.UserRepository
}

func NewProfileService(repo models.UserRepository) models.ProfileService {
	return &profileService{userRepository: repo}
}

func (ps *profileService) FetchProfileByID(id string) (*models.Profile, error) {
	user, err := ps.userRepository.FetchByID(id)
	if err != nil {
		return nil, err
	}
	return &models.Profile{FullName: user.FullName, Email: user.Email, Role: user.Role}, nil
}

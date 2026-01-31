package service

import (
	"api-integration/model"
	"api-integration/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserServices(repo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}

}

func (us *UserService) GetAllUsers() ([]model.User, error) {
	return us.userRepo.GetAllUsers()
}

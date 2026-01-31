package repository

import "api-integration/model"

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
}

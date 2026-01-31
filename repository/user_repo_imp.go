package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api-integration/model"
)

type userRepository struct {
	client *http.Client
	apiURL string
}

func NewUserRepository(client *http.Client, apiURL string) UserRepository {
	return &userRepository{
		client: client,
		apiURL: apiURL,
	}
}

func (ur *userRepository) GetAllUsers() ([]model.User, error) {
	resp, err := ur.client.Get(ur.apiURL)
	fmt.Println(resp)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponce model.UserAPIResponce
	err = json.NewDecoder(resp.Body).Decode(&apiResponce)
	if err != nil {
		return nil, err
	}
	// fmt.Println("-----", user)
	return apiResponce.Users, nil
}

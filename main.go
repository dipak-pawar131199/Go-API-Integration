package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"api-integration/config"
	"api-integration/handler"
	"api-integration/repository"
	"api-integration/service"
)

func main() {
	fmt.Println("Hello welcome to golang API integration")
	client := config.NewHTTPClient()

	userRepo := repository.NewUserRepository(client, config.APIURL)
	service := service.NewUserServices(userRepo)
	handler := handler.NewUserHandle(*service)

	r := gin.Default()

	r.POST("api/users", handler.GetUsers)

	r.Run(":8080")
}

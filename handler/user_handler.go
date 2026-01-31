package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api-integration/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandle(ser service.UserService) *UserHandler {
	return &UserHandler{service: ser}
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	user, err := h.service.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"user": user})
}

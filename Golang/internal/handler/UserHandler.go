package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamalakshancg/GolangPOC/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUserWithOrdersHandler(c *gin.Context) {
	res, err := h.UserService.GetUserWithOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

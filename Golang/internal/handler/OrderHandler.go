package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamalakshancg/GolangPOC/internal/service"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func (h *OrderHandler) GetOrdersHandler(c *gin.Context) {
	res, err := h.OrderService.GetOrderDetails()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

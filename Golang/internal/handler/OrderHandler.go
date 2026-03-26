package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamalakshancg/GolangPOC/internal/service"
)

type OrderHandler struct {
	OrderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

func (h *OrderHandler) Test2(c *gin.Context) {
	res, err := h.OrderService.ExecuteTest2()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *OrderHandler) Test1(c *gin.Context) {
	res, err := h.OrderService.ExecuteTest1()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, res)
}

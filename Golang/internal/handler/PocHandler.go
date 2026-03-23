package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kamalakshancg/GolangPOC/internal/service"
)

type PocHandler struct {
	Service *service.PocService
}

func NewPocHandler(service *service.PocService) *PocHandler {
	return &PocHandler{Service: service}
}

func (h *PocHandler) Test1(c *gin.Context) {
	res, err := h.Service.ExecuteTest1()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, res)
}

func (h *PocHandler) Test2(c *gin.Context) {
	res, err := h.Service.ExecuteTest2()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *PocHandler) Test3(c *gin.Context) {
	res, err := h.Service.ExecuteTest3()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

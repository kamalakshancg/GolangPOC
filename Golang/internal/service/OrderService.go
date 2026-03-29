package service

import (
	"github.com/kamalakshancg/GolangPOC/internal/entity"
	"github.com/kamalakshancg/GolangPOC/internal/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepo
}

func (s *OrderService) GetOrderDetails() ([]entity.Order, error) {
	return s.OrderRepo.GetOrders()
}

package service

import (
	"github.com/kamalakshancg/GolangPOC/internal/entity"
	"github.com/kamalakshancg/GolangPOC/internal/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepo
}

func NewOrderService(orderRepo *repository.OrderRepo) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

func (s *OrderService) GetOrderDetails() ([]entity.Order, error) {
	return s.OrderRepo.GetOrders()
}

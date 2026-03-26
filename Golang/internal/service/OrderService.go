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

func (s *OrderService) ExecuteTest2() ([]entity.Order, error) {
	return s.OrderRepo.GetWideOrders()
}

func (s *OrderService) ExecuteTest1() (string, error) {
	return s.OrderRepo.PingDB()
}

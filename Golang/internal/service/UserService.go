package service

import (
	"github.com/kamalakshancg/GolangPOC/internal/entity"
	"github.com/kamalakshancg/GolangPOC/internal/repository"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func (s *UserService) GetUserWithOrders() ([]entity.User, error) {
	rows, err := s.UserRepo.GetUserWithOrders()
	if err != nil {
		return nil, err
	}

	// Use a slice to preserve order, and a map to track slice indices
	var result []entity.User
	userIndexMap := make(map[int]int)

	for _, row := range rows {
		// 1. Find or Create User
		idx, exists := userIndexMap[row.UserID]
		if !exists {
			newUser := entity.User{
				ID:     row.UserID,
				Name:   row.UserName,
				Email:  row.Email,
				Orders: []entity.Order{},
			}
			result = append(result, newUser)
			idx = len(result) - 1
			userIndexMap[row.UserID] = idx
		}

		// 2. Find or Create Order
		var currentOrder *entity.Order
		for i := range result[idx].Orders {
			if result[idx].Orders[i].ID == row.OrderID {
				currentOrder = &result[idx].Orders[i]
				break
			}
		}

		if currentOrder == nil {
			newOrder := entity.Order{
				ID:          row.OrderID,
				UserID:      row.UserID,
				Amount:      row.Amount,
				Status:      row.Status,
				Description: row.Description,
				Items:       []entity.Item{},
			}
			result[idx].Orders = append(result[idx].Orders, newOrder)
			currentOrder = &result[idx].Orders[len(result[idx].Orders)-1]
		}

		// 3. Add Item
		item := entity.Item{
			ID:          row.ItemID,
			ProductName: row.Product,
			Quantity:    row.Quantity,
			Price:       row.Price,
		}
		currentOrder.Items = append(currentOrder.Items, item)
	}
	return result, nil
}

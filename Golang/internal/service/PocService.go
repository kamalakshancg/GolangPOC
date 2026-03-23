package service

import (
	"github.com/kamalakshancg/GolangPOC/internal/entity"
	"github.com/kamalakshancg/GolangPOC/internal/repository"
)

type PocService struct {
	Repo *repository.PocRepo
}

func NewPocService(repo *repository.PocRepo) *PocService {
	return &PocService{Repo: repo}
}

func (s *PocService) ExecuteTest1() (string, error) {
	return s.Repo.PingDB()
}

func (s *PocService) ExecuteTest2() ([]entity.Order, error) {
	return s.Repo.GetWideOrders()
}

func (s *PocService) ExecuteTest3() ([]entity.UserWithOrders, error) {
	rows, err := s.Repo.GetDeepNestedRows()
	if err != nil {
		return nil, err
	}

	// Manual mapping logic to group flat rows into nested JSON structure
	userMap := make(map[int]*entity.UserWithOrders)
	for _, row := range rows {
		if _, exists := userMap[row.UserID]; !exists {
			userMap[row.UserID] = &entity.UserWithOrders{ID: row.UserID, Name: row.UserName, Orders: []entity.NestedOrder{}}
		}

		var currentOrder *entity.NestedOrder
		for i, o := range userMap[row.UserID].Orders {
			if o.ID == row.OrderID {
				currentOrder = &userMap[row.UserID].Orders[i]
				break
			}
		}

		if currentOrder == nil {
			newOrder := entity.NestedOrder{ID: row.OrderID, Amount: row.Amount, Items: []entity.Item{}}
			userMap[row.UserID].Orders = append(userMap[row.UserID].Orders, newOrder)
			currentOrder = &userMap[row.UserID].Orders[len(userMap[row.UserID].Orders)-1]
		}

		item := entity.Item{ID: row.ItemID, ProductName: row.Product, Quantity: row.Quantity, Price: row.Price}
		currentOrder.Items = append(currentOrder.Items, item)
	}

	var result []entity.UserWithOrders
	for _, v := range userMap {
		result = append(result, *v)
	}
	return result, nil
}

package service

import (
	"fmt"
	"time"

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

func (s *PocService) ExecuteTest3() ([]entity.User, error) {
	// Start the stopwatch for your spreadsheet
	start := time.Now()

	rows, err := s.Repo.GetDeepNestedRows()
	if err != nil {
		return nil, err
	}

	// Capture exact DB Time
	dbFetchTime := time.Since(start).Milliseconds()
	fmt.Printf("Test 3 DB Fetch Time: %d ms\n", dbFetchTime)

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
				Amount:      row.Amount,
				Status:      row.Status,
				Description: row.Description,
				Items:       []entity.Item{},
			}
			result[idx].Orders = append(result[idx].Orders, newOrder)
			// Point to the newly appended order in the slice
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

	// Capture Total internal time
	totalInternalTime := time.Since(start).Milliseconds()
	fmt.Printf("Test 3 Total Internal Time: %d ms\n", totalInternalTime)

	return result, nil
}

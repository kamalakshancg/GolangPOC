package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kamalakshancg/GolangPOC/internal/entity"
)

type OrderRepo struct {
	Db *sqlx.DB
}

func (r *OrderRepo) GetOrders() ([]entity.Order, error) {
	var orders []entity.Order
	query := "SELECT id, user_id, amount, status, description FROM orders order by id LIMIT 1000"
	err := r.Db.Select(&orders, query)
	return orders, err
}

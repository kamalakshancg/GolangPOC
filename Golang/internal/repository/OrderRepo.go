package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kamalakshancg/GolangPOC/internal/entity"
)

type OrderRepo struct {
	Db *sqlx.DB
}

func (r *OrderRepo) GetWideOrders() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.Db.Select(&orders, "SELECT id, user_id, amount, status, description FROM orders order by id LIMIT 1000")
	return orders, err
}

func (r *OrderRepo) PingDB() (string, error) {
	var status string
	err := r.Db.Get(&status, "SELECT 'pong'")
	return status, err
}

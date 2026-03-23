package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kamalakshancg/GolangPOC/internal/entity"
)

type PocRepo struct {
	Db *sqlx.DB
}

func (r *PocRepo) PingDB() (string, error) {
	var status string
	err := r.Db.Get(&status, "SELECT 'pong'")
	return status, err
}

func (r *PocRepo) GetWideOrders() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.Db.Select(&orders, "SELECT * FROM orders LIMIT 1000")
	return orders, err
}

func (r *PocRepo) GetDeepNestedRows() ([]entity.FlatJoinRow, error) {
	query := `
		SELECT u.id as u_id, u.name as u_name, o.id as o_id, o.amount as o_amount, 
		       i.id as i_id, i.product_name as i_product, i.quantity as i_qty, i.price as i_price
		FROM users u
		JOIN orders o ON u.id = o.user_id
		JOIN items i ON o.id = i.order_id
		WHERE u.id BETWEEN 1 AND 50`

	var rows []entity.FlatJoinRow
	err := r.Db.Select(&rows, query)
	return rows, err
}

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
	err := r.Db.Select(&orders, "SELECT * FROM orders order by id LIMIT 1000")
	return orders, err
}

func (r *PocRepo) GetDeepNestedRows() ([]entity.FlatJoinRow, error) {
	query := `
        SELECT u.id as u_id,u.name as u_name, o.id as o_id,  
               o.amount as o_amount, i.id as i_id,  
               i.product_name as i_product, i.quantity as i_qty, i.price as i_price,
			   u.email as email, o.status as status, o.description as description
        FROM users u 
        INNER JOIN orders o ON u.id = o.user_id
        INNER JOIN items i ON o.id = i.order_id
        WHERE o.amount > 500 AND o.status = 'COMPLETED'
        ORDER BY u.id 
        LIMIT 1000`

	var rows []entity.FlatJoinRow
	err := r.Db.Select(&rows, query)
	return rows, err
}

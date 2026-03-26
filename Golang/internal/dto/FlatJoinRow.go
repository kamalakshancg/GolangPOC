package dto

type FlatJoinRow struct {
	UserID      int     `db:"u_id"`
	UserName    string  `db:"u_name"`
	OrderID     int     `db:"o_id"`
	Amount      float64 `db:"o_amount"`
	ItemID      int     `db:"i_id"`
	Product     string  `db:"i_product"`
	Quantity    int     `db:"i_qty"`
	Price       float64 `db:"i_price"`
	Email       string  `db:"email"`
	Status      string  `db:"status"`
	Description string  `db:"description"`
}

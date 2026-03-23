package entity

type Order struct {
	ID          int     `db:"id" json:"id"`
	UserID      int     `db:"user_id" json:"user_id"`
	Amount      float64 `db:"amount" json:"amount"`
	Status      string  `db:"status" json:"status"`
	Description string  `db:"description" json:"description"`
}

type Item struct {
	ID          int     `db:"id" json:"id"`
	OrderID     int     `db:"order_id" json:"order_id"`
	ProductName string  `db:"product_name" json:"product_name"`
	Quantity    int     `db:"quantity" json:"quantity"`
	Price       float64 `db:"price" json:"price"`
}

type UserWithOrders struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Orders []NestedOrder `json:"orders"`
}

type NestedOrder struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Items  []Item  `json:"items"`
}

type FlatJoinRow struct {
	UserID   int     `db:"u_id"`
	UserName string  `db:"u_name"`
	OrderID  int     `db:"o_id"`
	Amount   float64 `db:"o_amount"`
	ItemID   int     `db:"i_id"`
	Product  string  `db:"i_product"`
	Quantity int     `db:"i_qty"`
	Price    float64 `db:"i_price"`
}

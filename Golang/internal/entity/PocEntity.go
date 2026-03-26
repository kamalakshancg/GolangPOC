package entity

// --- Core Domain Models ---

type User struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email,omitempty"`
	Orders []Order `json:"orders"` // Slice of actual Orders
}

type Order struct {
	ID          int     `json:"id"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	Description string  `json:"description"`
	Items       []Item  `json:"items"` // Slice of actual Items
}

type Item struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

// --- Database Projection Model ---

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

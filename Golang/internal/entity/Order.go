package entity

type Order struct {
	ID          int     `db:"id" json:"id"`
	UserID      int     `db:"user_id" json:"user_id"`
	Amount      float64 `db:"amount" json:"amount"`
	Status      string  `db:"status" json:"status"`
	Description string  `db:"description" json:"description"`
	Items       []Item  `json:"items"`
}

package entity

type Order struct {
	ID          int     `db:"id" json:"id"`
	UserID      int     `db:"user_id" json:"user_id"` // This is the exact fix!
	Amount      float64 `db:"amount" json:"amount"`
	Status      string  `db:"status" json:"status"`
	Description string  `db:"description" json:"description"`
	Items       []Item  `json:"items"` // sqlx ignores this since it's not in the SELECT, which is perfect
}

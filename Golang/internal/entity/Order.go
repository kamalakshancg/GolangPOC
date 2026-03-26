package entity

type Order struct {
	ID          int     `json:"id"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	Description string  `json:"description"`
	Items       []Item  `json:"items"` // Slice of actual Items
}

package models

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Description string  `json:"description"`
}

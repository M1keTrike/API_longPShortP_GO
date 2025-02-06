package entities

type PriceChange struct {
	ProductID int     `json:"product_id"`
	OldPrice  float64 `json:"old_price"`
	NewPrice  float64 `json:"new_price"`
	ChangeTime string  `json:"change-time"`
	Visited   bool    `json:"visited"`
}
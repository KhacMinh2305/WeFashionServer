package cart

type CartRequestUpadte struct {
	Sku    int    `json:"sku"`
	Change string `json:"change"`
	Amount int    `json:"amount"`
}

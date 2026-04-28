package order

import "time"

type orderWithProductAmountRaw struct {
	Id            int       `json:"id"`
	Discount      float32   `json:"discount"`
	ShippingFee   float32   `json:"shipping_fee"`
	TotalPrice    float32   `json:"total_price"`
	OrderState    int       `json:"order_state"`
	ShippingState int       `json:"shipping_state"`
	CreatedAt     time.Time `json:"created_at"`
	UserId        int       `json:"user_id"`
	AddressId     int       `json:"address_id"`
	PaymentId     int       `json:"payment_id"`
	ShipperId     int       `json:"shipper_id"`
	ProductAmount int       `json:"product_amount"`
}

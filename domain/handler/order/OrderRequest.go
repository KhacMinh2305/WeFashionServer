package order

import (
	"time"
)

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

type OrderItem struct {
	Sku       int     `json:"sku"`
	Amount    int     `json:"amount"`
	Price     float32 `json:"price"`
	ProductId int     `json:"product_id"`
	SizeId    int     `json:"size_id"`
	ColorId   int     `json:"color_id"`
}

type OrderRequestBody struct {
	Discount    float32     `json:"discount"`
	ShippingFee float32     `json:"shipping_fee"`
	TotalPrice  float32     `json:"total_price"`
	UserId      int         `json:"user_id"`
	AddressId   int         `json:"address_id"`
	Items       []OrderItem `json:"items"`
}

// CreatePaymentLinkRequest represents the request to create a payment link
// type CreatePaymentLinkRequest struct {
// 	OrderCode        int64             `json:"orderCode"`
// 	Amount           int               `json:"amount"`
// 	Description      string            `json:"description"`
// 	CancelUrl        string            `json:"cancelUrl"`
// 	ReturnUrl        string            `json:"returnUrl"`
// 	Signature        *string           `json:"signature,omitempty"`
// 	Items            []PaymentLinkItem `json:"items,omitempty"`
// 	BuyerName        *string           `json:"buyerName,omitempty"`
// 	BuyerCompanyName *string           `json:"buyerCompanyName,omitempty"`
// 	BuyerTaxCode     *string           `json:"buyerTaxCode,omitempty"`
// 	BuyerEmail       *string           `json:"buyerEmail,omitempty"`
// 	BuyerPhone       *string           `json:"buyerPhone,omitempty"`
// 	BuyerAddress     *string           `json:"buyerAddress,omitempty"`
// 	Invoice          *InvoiceRequest   `json:"invoice,omitempty"`
// 	ExpiredAt        *int              `json:"expiredAt,omitempty"`
// }

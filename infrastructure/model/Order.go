package model

import "time"

type Order struct {
	Id            int       `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Discount      float32   `gorm:"not null" json:"discount"`
	ShippingFee   float32   `gorm:"not null" json:"shipping_fee"`
	TotalPrice    float32   `gorm:"not null" json:"total_price"`
	OrderState    int       `gorm:"not null;index" json:"order_state"`
	ShippingState int       `gorm:"not null;index" json:"shipping_state"`
	CreatedAt     time.Time `gorm:"not null" json:"created_at"`
	UserId        int       `gorm:"not null;index" json:"user_id"`
	AddressId     int       `gorm:"not null;index" json:"address_id"`
	PaymentId     int       `gorm:"not null;index" json:"payment_id"`
	ShipperId     int       `gorm:"not null;index" json:"shipper_id"`
}

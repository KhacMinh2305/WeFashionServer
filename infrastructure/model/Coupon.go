package model

import "time"

type Coupon struct {
	Id            int       `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name          string    `gorm:"not null;index" json:"name"`
	BannerUrl     string    `gorm:"not null" json:"banner_url"`
	DiscountValue float32   `gorm:"not null" json:"discount_value"`
	DiscountType  string    `gorm:"not null;index" json:"discount_type"`
	Amount        int       `gorm:"not null" json:"amount"`
	CreatedAt     time.Time `gorm:"not null" json:"created_at"`
	ExpiredAt     time.Time `gorm:"not null" json:"expired_at"`
	MaxDiscount   float32   `gorm:"not null" json:"max_discount"`
	MinOrderValue float32   `gorm:"not null" json:"min_order_value"`
	ShopId        int       `gorm:"not null;index" json:"shop_id"`
}

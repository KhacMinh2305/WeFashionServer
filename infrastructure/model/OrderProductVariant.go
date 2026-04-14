package model

type OrderProductVariant struct {
	OrderId int     `gorm:"not null;index;primaryKey" json:"order_id"`
	Sku     int     `gorm:"not null;index;primaryKey" json:"sku"`
	Amount  int     `gorm:"not null" json:"amount"`
	Price   float32 `gorm:"not null" json:"price"`
}

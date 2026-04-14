package model

type ProductVariant struct {
	Sku       int     `gorm:"primaryKey;autoIncrement;not null" json:"sku"`
	Amount    int     `gorm:"not null" json:"amount"`
	Price     float32 `gorm:"not null" json:"price"`
	ProductId int     `gorm:"not null;index" json:"product_id"`
	SizeId    int     `gorm:"not null;index" json:"size_id"`
	ColorId   int     `gorm:"not null;index" json:"color_id"`
}

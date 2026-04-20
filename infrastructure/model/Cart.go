package model

type Cart struct {
	UserId int `gorm:"not null;index;primaryKey" json:"user_id"`
	SkuId  int `gorm:"not null;index;primaryKey" json:"sku_id"`
	Amount int `gorm:"not null" json:"amount"`
}

package model

import "time"

type FavoriteProduct struct {
	ProductId int       `gorm:"not null;index;primaryKey" json:"product_id"`
	UserId    int       `gorm:"not null;index;primaryKey" json:"user_id"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

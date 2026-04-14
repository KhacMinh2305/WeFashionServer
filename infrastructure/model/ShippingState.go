package model

import "time"

type ShippingState struct {
	Id        int       `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Detail    string    `gorm:"not null" json:"detail"`
	Location  string    `gorm:"not null" json:"location"`
	Latitude  float64   `gorm:"not null" json:"latitude"`
	Longitude float64   `gorm:"not null" json:"longitude"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	OrderId   int       `gorm:"not null;index" json:"order_id"`
}

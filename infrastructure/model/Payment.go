package model

import "time"

type Payment struct {
	Id          int       `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	Description string    `gorm:"not null" json:"description"`
	OrderId     int       `gorm:"not null;index" json:"order_id"`
	UserId      int       `gorm:"not null;index" json:"user_id"`
}

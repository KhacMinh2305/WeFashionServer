package model

type Address struct {
	Id           int     `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name         string  `gorm:"not null" json:"name"`
	Ward         string  `gorm:"not null" json:"ward"`
	District     string  `gorm:"not null" json:"district"`
	City         string  `gorm:"not null" json:"city"`
	Detail       string  `gorm:"not null" json:"detail"`
	Latitude     float64 `gorm:"not null" json:"latitude"`
	Longitude    float64 `gorm:"not null" json:"longitude"`
	ReceiverName string  `gorm:"not null" json:"receiver_name"`
	Phone        string  `gorm:"not null;index" json:"phone"`
	IsDefault    bool    `gorm:"not null;index" json:"is_default"`
	UserId       int     `gorm:"not null;index" json:"user_id"`
}

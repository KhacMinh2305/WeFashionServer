package model

type Shipper struct {
	Id          int    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name        string `gorm:"not null;index" json:"name"`
	ImageUrl    string `gorm:"not null" json:"image_url"`
	PhoneNumber string `gorm:"not null;index" json:"phone_number"`
}

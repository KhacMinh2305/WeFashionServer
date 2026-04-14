package model

type User struct {
	Id          int    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name        string `gorm:"not null;index" json:"name"`
	AvatarUrl   string `gorm:"not null" json:"avatar_url"`
	Email       string `gorm:"not null;index" json:"email"`
	PhoneNumber string `gorm:"not null;index" json:"phone_number"`
	Bio         string `gorm:"not null" json:"bio"`
}

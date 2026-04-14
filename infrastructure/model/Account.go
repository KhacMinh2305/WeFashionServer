package model

type Account struct {
	Id       int    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Username string `gorm:"not null;index" json:"username"`
	Password string `gorm:"not null" json:"password"`
	UserId   int    `gorm:"not null;index" json:"user_id"`
}

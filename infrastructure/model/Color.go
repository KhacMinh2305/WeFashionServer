package model

type Color struct {
	Id  int    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Rgb string `gorm:"not null;index" json:"rgb"`
}

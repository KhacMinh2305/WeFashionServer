package model

type Category struct {
	Id   int    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name string `gorm:"not null;index" json:"name"`
}

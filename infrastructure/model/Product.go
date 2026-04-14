package model

type Product struct {
	Id          int     `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name        string  `gorm:"not null;index" json:"name"`
	ImageUrl    string  `gorm:"not null" json:"image_url"`
	Description string  `gorm:"not null" json:"description"`
	Rating      float32 `gorm:"not null;index" json:"rating"`
	SoldAmount  int     `gorm:"not null;index" json:"sold_amount"`
	LikedAmount int     `gorm:"not null;index" json:"liked_amount"`
	CategoryId  int     `gorm:"not null;index" json:"category_id"`
	ShopId      int     `gorm:"not null;index" json:"shop_id"`
}

package search

type SearchedProduct struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"image_url"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	SoldAmount  int     `json:"sold_amount"`
	LikedAmount int     `json:"liked_amount"`
	CategoryId  int     `json:"category_id"`
	ShopId      int     `json:"shop_id"`
}

type SearchedShop struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	AvatarUrl   string  `json:"avatar_url"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	Bio         string  `json:"bio"`
	RateAmount  int     `json:"rate_amount"`
	Rating      float32 `json:"rating"`
	Followers   int     `json:"followers"`
}

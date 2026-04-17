package shop

type ShopResponse struct {
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

type ShopListResponse struct {
	Shops []ShopResponse `json:"shops"`
}

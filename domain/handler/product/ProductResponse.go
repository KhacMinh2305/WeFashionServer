package product

type ProductListResponse struct {
	Title  string         `json:"title"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
	Data   []ProductBrief `json:"data"`
}

type ProductBrief struct {
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

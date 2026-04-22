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

// ProductDetailResponse is the main response for product detail
type ProductDetailResponse struct {
	Product ProductDetail `json:"product"`
	Sku     []ProductVariantDetail `json:"sku"`
	Shop    ShopBrief `json:"shop"`
}

// ProductDetail omits ShopId
type ProductDetail struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"image_url"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	SoldAmount  int     `json:"sold_amount"`
	LikedAmount int     `json:"liked_amount"`
	CategoryId  int     `json:"category_id"`
}

type ProductVariantDetail struct {
	Sku    int         `json:"sku"`
	Amount int         `json:"amount"`
	Price  float32     `json:"price"`
	Size   SizeBrief   `json:"size"`
	Color  ColorBrief  `json:"color"`
}

type SizeBrief struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ColorBrief struct {
	Id  int    `json:"id"`
	Rgb string `json:"rgb"`
}

type ShopBrief struct {
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

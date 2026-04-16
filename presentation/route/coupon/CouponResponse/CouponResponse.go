package CouponResponse

type CouponResponse struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	BannerUrl     string  `json:"banner_url"`
	DiscountValue float32 `json:"discount_value"`
	DiscountType  string  `json:"discount_type"`
	Amount        int     `json:"amount"`
	CreatedAt     string  `json:"created_at"`
	ExpiredAt     string  `json:"expired_at"`
	MaxDiscount   float32 `json:"max_discount"`
	MinOrderValue float32 `json:"min_order_value"`
	ShopId        int     `json:"shop_id"`
}

type CouponListResponse struct {
	Coupons []CouponResponse `json:"coupons"`
}

type CouponShopListResponse struct {
	ShopId  int              `json:"shop_id"`
	Coupons []CouponResponse `json:"coupons"`
}

type CouponUserListResponse struct {
	UserId  int              `json:"user_id"`
	Coupons []CouponResponse `json:"coupons"`
}

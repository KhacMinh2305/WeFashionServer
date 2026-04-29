package order

import "time"

type OrderWithProductAmountResponse struct {
	Id            int       `json:"id"`
	Discount      float32   `json:"discount"`
	ShippingFee   float32   `json:"shipping_fee"`
	TotalPrice    float32   `json:"total_price"`
	OrderState    int       `json:"order_state"`
	ShippingState int       `json:"shipping_state"`
	CreatedAt     time.Time `json:"created_at"`
	UserId        int       `json:"user_id"`
	AddressId     int       `json:"address_id"`
	PaymentId     int       `json:"payment_id"`
	ShipperId     int       `json:"shipper_id"`
	ProductAmount int       `json:"product_amount"`
}

type OrderListByUserResponse struct {
	Orders []OrderWithProductAmountResponse `json:"orders"`
}

type OrderPaymentLinkResponse struct {
	CheckoutUrl string `json:"checkout_url"`
	QrCode      string `json:"qr_code"`
}

type OrderDetailOrderResponse struct {
	Id          int       `json:"id"`
	Discount    float32   `json:"discount"`
	ShippingFee float32   `json:"shipping_fee"`
	TotalPrice  float32   `json:"total_price"`
	OrderState  int       `json:"order_state"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderDetailAddressResponse struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Ward         string  `json:"ward"`
	District     string  `json:"district"`
	City         string  `json:"city"`
	Detail       string  `json:"detail"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	ReceiverName string  `json:"receiver_name"`
	Phone        string  `json:"phone"`
	IsDefault    bool    `json:"is_default"`
}

type OrderDetailPaymentResponse struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	OrderId     int       `json:"order_id"`
}

type OrderDetailCategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OrderDetailShopResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OrderDetailSizeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OrderDetailColorResponse struct {
	Id  int    `json:"id"`
	Rgb string `json:"rgb"`
}

type OrderDetailProductSkuResponse struct {
	Sku    int                      `json:"sku"`
	Amount int                      `json:"amount"`
	Price  float32                  `json:"price"`
	Size   OrderDetailSizeResponse  `json:"size"`
	Color  OrderDetailColorResponse `json:"color"`
}

type OrderDetailProductResponse struct {
	Id          int                             `json:"id"`
	Name        string                          `json:"name"`
	ImageUrl    string                          `json:"image_url"`
	Description string                          `json:"description"`
	Rating      float32                         `json:"rating"`
	SoldAmount  int                             `json:"sold_amount"`
	LikedAmount int                             `json:"liked_amount"`
	Category    OrderDetailCategoryResponse     `json:"category"`
	Shop        OrderDetailShopResponse         `json:"shop"`
	Skus        []OrderDetailProductSkuResponse `json:"skus"`
}

type OrderDetailResponse struct {
	Order          OrderDetailOrderResponse     `json:"order"`
	User           OrderUserResponse            `json:"user"`
	ShippingStates []OrderShippingStateResponse `json:"shipping_states"`
	Address        OrderDetailAddressResponse   `json:"address"`
	Payment        OrderDetailPaymentResponse   `json:"payment"`
	Shipper        OrderShipperResponse         `json:"shipper"`
	Products       []OrderDetailProductResponse `json:"products"`
}

type OrderUserResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	AvatarUrl   string `json:"avatar_url"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Bio         string `json:"bio"`
}

type OrderShipperResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ImageUrl    string `json:"image_url"`
	PhoneNumber string `json:"phone_number"`
}

type OrderShippingStateResponse struct {
	Id        int       `json:"id"`
	Detail    string    `json:"detail"`
	Location  string    `json:"location"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	OrderId   int       `json:"order_id"`
}

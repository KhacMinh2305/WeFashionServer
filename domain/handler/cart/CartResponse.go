package cart

type CartRawItem struct {
	SkuId  int
	Amount int
}

type CartUser struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	AvatarUrl   string `json:"avatar_url"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Bio         string `json:"bio"`
}

type CartCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CartProductColor struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CartProductSize struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CartProductSku struct {
	Sku       int              `json:"sku"`
	Amount    int              `json:"amount"`
	Price     float32          `json:"price"`
	ProductId int              `json:"product_id"`
	Size      CartProductSize  `json:"size"`
	Color     CartProductColor `json:"color"`
}

type CartProduct struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	ImageUrl    string            `json:"image_url"`
	Description string            `json:"description"`
	Rating      float32           `json:"rating"`
	SoldAmount  int               `json:"sold_amount"`
	LikedAmount int               `json:"liked_amount"`
	Category    CartCategory      `json:"category"`
	Skus        *[]CartProductSku `json:"skus"`
}

type CartResponse struct {
	User     CartUser       `json:"user"`
	Products *[]CartProduct `json:"products"`
}

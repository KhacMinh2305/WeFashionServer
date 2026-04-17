package shop

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/handler/shop"
)

func RegisterShopRoutes() {
	di.Router.GET("api/shop", shop.GetShops)
	di.Router.PUT("api/shop/follow", shop.UpdateShopFollow)
}

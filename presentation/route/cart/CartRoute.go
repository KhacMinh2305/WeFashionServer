package cart

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/handler/cart"
)

func RegisterCartRoute() {
	di.Router.GET("api/cart/:id", cart.GetUserCart)
	di.Router.POST("api/cart/:id/sku/update", cart.UpdateSku)
	di.Router.DELETE("api/cart/:id/delete", cart.RemoveFromCart)
}

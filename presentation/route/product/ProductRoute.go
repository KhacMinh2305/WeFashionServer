package product

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/handler/product"
)

func RegisterProductRoute() {
	di.Router.GET("api/product", product.GetProducts)
	di.Router.GET("api/product/top-rated", product.GetTopRatedProducts)
	di.Router.GET("api/product/best-seller", product.GetBestSellerProducts)
	di.Router.GET("api/product/most-liked", product.GetMostLikedProducts)
	di.Router.GET("api/product/:id/details", product.GetProductDetail)

	// variants
	di.Router.GET("api/product/shop/:id", product.GetShopProducts)
}

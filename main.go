package main

import (
	"WeFashionServer/di"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/presentation/route/authentication"
	"WeFashionServer/presentation/route/category"
	"WeFashionServer/presentation/route/color"
	"WeFashionServer/presentation/route/coupon"
	"WeFashionServer/presentation/route/shop"
)

func main() {

	database.Connect()

	registerRoutes()

	di.Router.Run()
}

func registerRoutes() {
	authentication.RegisterAuthenticationRoute()
	color.RegisterColorRoutes()
	category.RegisterCategoryRoutes()
	shop.RegisterShopRoutes()
	coupon.RegisterCouponRoutes()
}

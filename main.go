package main

import (
	"WeFashionServer/di"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/presentation/route/account"
	"WeFashionServer/presentation/route/address"
	"WeFashionServer/presentation/route/authentication"
	"WeFashionServer/presentation/route/category"
	"WeFashionServer/presentation/route/color"
	"WeFashionServer/presentation/route/coupon"
	"WeFashionServer/presentation/route/shop"
	"WeFashionServer/presentation/route/user"
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

	// important
	account.RegisterAccountRoutes()
	user.RegisterUserRoutes()
	address.RegisterAddressRoutes()
}

package main

import (
	"WeFashionServer/di"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/presentation/route/authentication"
	"WeFashionServer/presentation/route/coupon"
)

func main() {

	database.Connect()

	registerRoutes()

	di.Router.Run()
}

func registerRoutes() {
	authentication.RegisterAuthenticationRoute()
	coupon.RegisterCouponRoutes()
}

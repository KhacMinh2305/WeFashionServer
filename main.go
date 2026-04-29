package main

import (
	"WeFashionServer/di"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/presentation/route/account"
	"WeFashionServer/presentation/route/address"
	"WeFashionServer/presentation/route/authentication"
	"WeFashionServer/presentation/route/cart"
	"WeFashionServer/presentation/route/category"
	"WeFashionServer/presentation/route/color"
	"WeFashionServer/presentation/route/coupon"
	"WeFashionServer/presentation/route/order"
	"WeFashionServer/presentation/route/product"
	"WeFashionServer/presentation/route/search"
	"WeFashionServer/presentation/route/shop"
	"WeFashionServer/presentation/route/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if dbErr := database.Connect(); dbErr != nil {
		log.Fatalf("Failed to connect to database: %v", dbErr)
		return
	}

	registerRoutes()

	di.PaymentRepo.Initialize()

	if err := di.Router.Run("0.0.0.0:" + port); err != nil {
		log.Fatalf("Failed to start server for port: %v", err)
	}
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

	product.RegisterProductRoute()
	search.RegisterSearchRoute()
	cart.RegisterCartRoute()
	order.RegisterOrderRoutes()
}

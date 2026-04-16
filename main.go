package main

import (
	"WeFashionServer/di"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/presentation/route/coupon"
	// "WeFashionServer/mock"
	// "fmt"
	// "time"
)

func main() {

	database.Connect()

	// if err := mock.InsertMockUnderwearsProducts(); err != nil {
	// 	fmt.Println("InsertMockCategories error: " + err.Error())
	// } else {
	// 	fmt.Println("InsertMockUnderwearsProducts inserted!")
	// }

	// if err := mock.InsertMockUnderwearsVariants(); err != nil {
	// 	fmt.Println("InsertMockCategories error: " + err.Error())
	// } else {
	// 	fmt.Println("InsertMockUnderwearsVariants inserted!")
	// }

	coupon.RegisterCouponRoutes()

	di.Router.Run()
}

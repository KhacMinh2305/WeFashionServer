package main

import (
	"WeFashionServer/infrastructure/database"
	// "WeFashionServer/mock"
	// "fmt"
	// "time"
)

func main() {

	database.Connect()

	//Chèn mock data sau 10 giây bằng goroutine
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	if err := mock.InsertMockColors(); err != nil {
	// 		fmt.Println("InsertMockCategories error: " + err.Error())
	// 	} else {
	// 		fmt.Println("Mock categories inserted!")
	// 	}
	// 	time.Sleep(5 * time.Second)
	// 	if err := mock.InsertMockSizes(); err != nil {
	// 		fmt.Println("InsertMockCategories error: " + err.Error())
	// 	} else {
	// 		fmt.Println("Mock categories inserted!")
	// 	}
	// }()

	// di.Router.Run()
}

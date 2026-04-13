package main

import (
	"WeFashionServer/di"
	"WeFashionServer/presentation/route"
)

func main() {
	route.RegisterAuthenticationRoute()

	di.Router.Run()
}

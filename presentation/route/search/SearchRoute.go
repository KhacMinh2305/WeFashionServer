package search

import "WeFashionServer/di"

func RegisterSearchRoute() {
	di.Router.GET("api/search")
}

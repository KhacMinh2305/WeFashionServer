package search

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/handler/search"
)

func RegisterSearchRoute() {
	di.Router.GET("api/search", search.SearchProductAndShop)
}

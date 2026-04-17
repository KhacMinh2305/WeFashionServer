package category

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/handler/category"
)

func RegisterCategoryRoutes() {
	// Route: /api/category (all categories or by id)
	di.Router.GET("api/category", category.GetCategories)
}

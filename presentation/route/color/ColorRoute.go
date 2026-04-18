package color

import (
	"WeFashionServer/di"
	colorHandler "WeFashionServer/domain/handler/color"
)

// GET /api/color?color_id=...
func RegisterColorRoutes() {
	di.Router.GET("api/color", colorHandler.GetColors)
}

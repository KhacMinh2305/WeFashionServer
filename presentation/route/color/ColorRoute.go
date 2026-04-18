
package color

import (
	"WeFashionServer/domain/handler/color"
	"github.com/gin-gonic/gin"
)

// GET /api/color?color_id=...
func RegisterColorRoutes() {
	router.GET("api/color", color.GetColors)
}


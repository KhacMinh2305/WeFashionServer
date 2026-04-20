package user

import (
	"WeFashionServer/di"
	userHandler "WeFashionServer/domain/handler/user"
)

func RegisterUserRoutes() {
	di.Router.GET("api/user", userHandler.GetUser)
	di.Router.POST("api/user/:id/update", userHandler.UpdateUser)
}

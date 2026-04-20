package address

import (
	"WeFashionServer/di"
	addressHandler "WeFashionServer/domain/handler/address"
)

func RegisterAddressRoutes() {
	di.Router.GET("api/address/user/:id", addressHandler.GetAddressesByUserId)
	di.Router.GET("api/address/:id", addressHandler.GetAddressById)
	di.Router.POST("api/address/create", addressHandler.CreateAddress)
	di.Router.PUT("api/address/update", addressHandler.UpdateAddress)
	di.Router.DELETE("api/address/:id/delete", addressHandler.DeleteAddress)
}

package account

import (
	"WeFashionServer/di"
	accountHandler "WeFashionServer/domain/handler/account"
)

func RegisterAccountRoutes() {
	di.Router.POST("api/account/register", accountHandler.RegisterAccount)
	di.Router.POST("api/account/login", accountHandler.LoginAccount)
	di.Router.POST("api/account/forgot-password", accountHandler.ForgotPassword)
	di.Router.POST("api/account/forgot-password/validate", accountHandler.ValidateCode)
}

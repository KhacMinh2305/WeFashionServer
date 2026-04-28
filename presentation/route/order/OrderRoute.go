package order

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/handler/order"
)

func RegisterOrderRoutes() {
	di.Router.GET("api/order/user/:id", order.GetUserOrders)
	di.Router.GET("api/order/:id/details", order.GetOrderDetails)
	// webhooks
	di.Router.POST("/api/order/payment/webhook", order.HandlePaymentWebhook)
	di.Router.GET("/api/order/payment/webhook", order.HandlePaymentWebhook)
}

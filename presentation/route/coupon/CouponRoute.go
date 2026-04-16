package coupon

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/service"
)

func RegisterCouponRoutes() {
	di.Router.GET("api/coupons", service.GetCoupons)
	di.Router.GET("api/coupons/:id", service.GetCouponsById)
	di.Router.GET("api/coupons/shop/:shop_id", service.GetCouponsOfShopByShopId)
	di.Router.GET("api/coupons/user/:user_id", service.GetAvailableCouponsForUserByUserId)
	di.Router.GET("api/coupons/order/shop/:shop_id", service.GetCouponsForOrderOfShopByShopId)
}

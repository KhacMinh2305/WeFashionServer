package coupon

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/service"
)

func RegisterCouponRoutes() {
	// Route: /api/coupons (all coupons)
	di.Router.GET("api/coupons", service.GetCoupons)
	// Route: /api/coupons/:id (get coupon by id)
	di.Router.GET("api/coupons/:id", service.GetCouponById)
	// Route: /api/coupons/shop (get coupons by shop_id)
	di.Router.GET("api/coupons/shop", service.GetCouponsOfShopByShopId)
	// Route: /api/coupons/user (get coupons by user_id)
	di.Router.GET("api/coupons/user", service.GetAvailableCouponsForUserByUserId)
	// Route: /api/coupons/order (get coupons for order by shop_id)
	di.Router.GET("api/coupons/order", service.GetCouponsForOrderOfShopByShopId)
}

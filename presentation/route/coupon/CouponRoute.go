package coupon

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/handler/coupon"
)

func RegisterCouponRoutes() {
	// Route: /api/coupons (all coupons)
	di.Router.GET("api/coupons", coupon.GetCoupons)
	// Route: /api/coupons/:id (get coupon by id)
	di.Router.GET("api/coupons/:id", coupon.GetCouponById)
	// Route: /api/coupons/shop (get coupons by shop_id)
	di.Router.GET("api/coupons/shop", coupon.GetCouponsOfShopByShopId)
	// Route: /api/coupons/user (get coupons by user_id)
	di.Router.GET("api/coupons/user", coupon.GetAvailableCouponsForUserByUserId)
	// Route: /api/coupons/order (get coupons for order by shop_id)
	di.Router.GET("api/coupons/order", coupon.GetCouponsForOrderOfShopByShopId)
}

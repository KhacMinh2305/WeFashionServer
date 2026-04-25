package coupon

import (
	"WeFashionServer/domain/entity"
	"WeFashionServer/domain/handler/authentication"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Route: /api/coupons (all coupons)
func GetCoupons(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	var coupons []model.Coupon
	if err := database.DB.Find(&coupons).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	resp := buildCouponResponseList(coupons)
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponListResponse{Coupons: resp},
	})
}

// Route: /api/coupons/:id
func GetCouponById(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	id, ok := getIntParam(ctx, "id", false)
	if !ok {
		return
	}
	var coupon model.Coupon
	if err := database.DB.First(&coupon, id).Error; err != nil {
		ctx.JSON(http.StatusOK, entity.SuccessReponse[*CouponResponse]{
			StatusCode: http.StatusOK,
			Time:       time.Now(),
			Data:       nil,
		})
		return
	}
	resp := buildCouponResponseList([]model.Coupon{coupon})
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       resp[0],
	})
}

// Route: /api/coupons/shop?shop_id=123
func GetCouponsOfShopByShopId(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	shopId, ok := getIntParam(ctx, "shop_id", true)
	if !ok {
		return
	}
	if !checkShopExists(ctx, shopId) {
		return
	}
	var coupons []model.Coupon
	if err := database.DB.Where("shop_id = ?", shopId).Find(&coupons).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	resp := buildCouponResponseList(coupons)
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponShopListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponShopListResponse{ShopId: shopId, Coupons: resp},
	})
}

// Route: /api/coupons/user?user_id=456
func GetAvailableCouponsForUserByUserId(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	userId, ok := getIntParam(ctx, "user_id", true)
	if !ok {
		return
	}
	if !checkUserExists(ctx, userId) {
		return
	}
	var coupons []model.Coupon
	if err := database.DB.Where("shop_id = ?", -1).Find(&coupons).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	resp := buildCouponResponseList(coupons)
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponUserListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponUserListResponse{UserId: userId, Coupons: resp},
	})
}

// Route: /api/coupons/order?shop_id=123
func GetCouponsForOrderOfShopByShopId(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	shopId, ok := getIntParam(ctx, "shop_id", true)
	if !ok {
		return
	}
	if !checkShopExists(ctx, shopId) {
		return
	}
	var coupons []model.Coupon
	if err := database.DB.Where("shop_id = ? OR shop_id = ?", -1, shopId).Find(&coupons).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	resp := buildCouponResponseList(coupons)
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponShopListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponShopListResponse{ShopId: shopId, Coupons: resp},
	})
}

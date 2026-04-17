import (
	"WeFashionServer/domain/service"
)
package service

import (
	"WeFashionServer/domain/entity"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"WeFashionServer/presentation/route/coupon/CouponResponse"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Helper: build coupon response list
func buildCouponResponseList(coupons []model.Coupon) []CouponResponse.CouponResponse {
	var resp []CouponResponse.CouponResponse
	for _, c := range coupons {
		resp = append(resp, CouponResponse.CouponResponse{
			Id:            c.Id,
			Name:          c.Name,
			BannerUrl:     c.BannerUrl,
			DiscountValue: c.DiscountValue,
			DiscountType:  c.DiscountType,
			Amount:        c.Amount,
			CreatedAt:     c.CreatedAt.Format(time.RFC3339),
			ExpiredAt:     c.ExpiredAt.Format(time.RFC3339),
			MaxDiscount:   c.MaxDiscount,
			MinOrderValue: c.MinOrderValue,
			ShopId:        c.ShopId,
		})
	}
	return resp
}


// Route: /api/coupons (all coupons)
func GetCoupons(ctx *gin.Context) {
	// Validate token
	_, err := service.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      "Unauthorized",
			Detail:     err.Error(),
		})
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
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse.CouponListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponResponse.CouponListResponse{Coupons: resp},
	})
}

// Route: /api/coupons/:id
func GetCouponById(ctx *gin.Context) {
	_, err := service.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      "Unauthorized",
			Detail:     err.Error(),
		})
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid id",
			Detail:     err.Error(),
		})
		return
	}
	var coupon model.Coupon
	if err := database.DB.First(&coupon, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Coupon not found",
			Detail:     err.Error(),
		})
		return
	}
	resp := buildCouponResponseList([]model.Coupon{coupon})
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse.CouponResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       resp[0],
	})
}

// Route: /api/coupons/shop?shop_id=123
func GetCouponsOfShopByShopId(ctx *gin.Context) {
	_, err := service.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      "Unauthorized",
			Detail:     err.Error(),
		})
		return
	}
	shopIdStr := ctx.Query("shop_id")
	if shopIdStr == "" {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Missing shop_id",
			Detail:     "shop_id query parameter is required",
		})
		return
	}
	shopId, err := strconv.Atoi(shopIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid shop_id",
			Detail:     err.Error(),
		})
		return
	}
	var shop model.Shop
	if err := database.DB.First(&shop, shopId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Shop not found",
			Detail:     err.Error(),
		})
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
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse.CouponShopListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponResponse.CouponShopListResponse{ShopId: shopId, Coupons: resp},
	})
}

// Route: /api/coupons/user?user_id=456
func GetAvailableCouponsForUserByUserId(ctx *gin.Context) {
	_, err := service.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      "Unauthorized",
			Detail:     err.Error(),
		})
		return
	}
	userIdStr := ctx.Query("user_id")
	if userIdStr == "" {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Missing user_id",
			Detail:     "user_id query parameter is required",
		})
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid user_id",
			Detail:     err.Error(),
		})
		return
	}
	var user model.User
	if err := database.DB.First(&user, userId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "User not found",
			Detail:     err.Error(),
		})
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
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse.CouponUserListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponResponse.CouponUserListResponse{UserId: userId, Coupons: resp},
	})
}

// Route: /api/coupons/order?shop_id=123
func GetCouponsForOrderOfShopByShopId(ctx *gin.Context) {
	_, err := service.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      "Unauthorized",
			Detail:     err.Error(),
		})
		return
	}
	shopIdStr := ctx.Query("shop_id")
	if shopIdStr == "" {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Missing shop_id",
			Detail:     "shop_id query parameter is required",
		})
		return
	}
	shopId, err := strconv.Atoi(shopIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid shop_id",
			Detail:     err.Error(),
		})
		return
	}
	var shop model.Shop
	if err := database.DB.First(&shop, shopId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Shop not found",
			Detail:     err.Error(),
		})
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
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse.CouponShopListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponResponse.CouponShopListResponse{ShopId: shopId, Coupons: resp},
	})
}


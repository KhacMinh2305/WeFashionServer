package coupon

import (
	"WeFashionServer/domain/entity"
	"WeFashionServer/domain/handler/authentication"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Helper: build coupon response list
func buildCouponResponseList(coupons []model.Coupon) []CouponResponse {
	var resp []CouponResponse
	for _, c := range coupons {
		resp = append(resp, CouponResponse{
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

// Helper: validate token, return true if valid, otherwise write error response and return false
func validateTokenOrAbort(ctx *gin.Context) bool {
	_, err := authentication.ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      "Unauthorized",
			Detail:     err.Error(),
		})
		return false
	}
	return true
}

// Helper: get int param from query or path, write error if invalid
func getIntParam(ctx *gin.Context, key string, fromQuery bool) (int, bool) {
	var valStr string
	if fromQuery {
		valStr = ctx.Query(key)
	} else {
		valStr = ctx.Param(key)
	}
	if valStr == "" {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Missing " + key,
			Detail:     key + " is required",
		})
		return 0, false
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid " + key,
			Detail:     err.Error(),
		})
		return 0, false
	}
	return val, true
}

// Helper: check existence of shop by id
func checkShopExists(ctx *gin.Context, shopId int) bool {
	var shop model.Shop
	if err := database.DB.First(&shop, shopId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Shop not found",
			Detail:     err.Error(),
		})
		return false
	}
	return true
}

// Helper: check existence of user by id
func checkUserExists(ctx *gin.Context, userId int) bool {
	var user model.User
	if err := database.DB.First(&user, userId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "User not found",
			Detail:     err.Error(),
		})
		return false
	}
	return true
}

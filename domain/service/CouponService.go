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

func GetCoupons(ctx *gin.Context) {
	var coupons []model.Coupon
	if err := database.DB.Find(&coupons).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
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
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse.CouponListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponResponse.CouponListResponse{Coupons: resp},
	})
}

func GetCouponsById(ctx *gin.Context) {
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
	resp := CouponResponse.CouponResponse{
		Id:            coupon.Id,
		Name:          coupon.Name,
		BannerUrl:     coupon.BannerUrl,
		DiscountValue: coupon.DiscountValue,
		DiscountType:  coupon.DiscountType,
		Amount:        coupon.Amount,
		CreatedAt:     coupon.CreatedAt.Format(time.RFC3339),
		ExpiredAt:     coupon.ExpiredAt.Format(time.RFC3339),
		MaxDiscount:   coupon.MaxDiscount,
		MinOrderValue: coupon.MinOrderValue,
		ShopId:        coupon.ShopId,
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse.CouponResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       resp,
	})
}

func GetCouponsOfShopByShopId(ctx *gin.Context) {
	shopIdStr := ctx.Param("shop_id")
	shopId, err := strconv.Atoi(shopIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid shop_id",
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
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse.CouponShopListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponResponse.CouponShopListResponse{ShopId: shopId, Coupons: resp},
	})
}

func GetAvailableCouponsForUserByUserId(ctx *gin.Context) {
	userIdStr := ctx.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid user_id",
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
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse.CouponUserListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponResponse.CouponUserListResponse{UserId: userId, Coupons: resp},
	})
}

func GetCouponsForOrderOfShopByShopId(ctx *gin.Context) {
	shopIdStr := ctx.Param("shop_id")
	shopId, err := strconv.Atoi(shopIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid shop_id",
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
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CouponResponse.CouponShopListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CouponResponse.CouponShopListResponse{ShopId: shopId, Coupons: resp},
	})
}

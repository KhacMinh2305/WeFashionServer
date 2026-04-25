package shop

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

// GET /api/shop : get all shops
func GetShops(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	var shops []model.Shop
	if err := database.DB.Find(&shops).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	resp := make([]ShopResponse, len(shops))
	for i, s := range shops {
		resp[i] = ShopResponse{
			Id:          s.Id,
			Name:        s.Name,
			AvatarUrl:   s.AvatarUrl,
			Email:       s.Email,
			PhoneNumber: s.PhoneNumber,
			Bio:         s.Bio,
			RateAmount:  s.RateAmount,
			Rating:      s.Rating,
			Followers:   s.Followers,
		}
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[ShopListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       ShopListResponse{Shops: resp},
	})
}

// PUT /api/shop/follow?shop_id=...&follow=...
func UpdateShopFollow(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	shopIdStr := ctx.Query("shop_id")
	followStr := ctx.Query("follow")
	if shopIdStr == "" || followStr == "" {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Missing parameter",
			Detail:     "shop_id and follow are required",
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
	follow, err := strconv.ParseBool(followStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid follow value",
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
	if follow {
		shop.Followers++
	} else {
		if shop.Followers == 0 {
			ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Error:      "Followers cannot be negative",
				Detail:     "Shop followers is already 0",
			})
			return
		}
		shop.Followers--
	}
	if err := database.DB.Model(&shop).Update("followers", shop.Followers).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Failed to update followers",
			Detail:     err.Error(),
		})
		return
	}
	resp := ShopResponse{
		Id:          shop.Id,
		Name:        shop.Name,
		AvatarUrl:   shop.AvatarUrl,
		Email:       shop.Email,
		PhoneNumber: shop.PhoneNumber,
		Bio:         shop.Bio,
		RateAmount:  shop.RateAmount,
		Rating:      shop.Rating,
		Followers:   shop.Followers,
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[ShopResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       resp,
	})
}

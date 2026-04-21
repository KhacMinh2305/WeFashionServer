package product

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

func parseLimitOffset(ctx *gin.Context) (int, int, *entity.ErrorResponse) {
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")
	limit, err1 := strconv.Atoi(limitStr)
	offset, err2 := strconv.Atoi(offsetStr)
	if err1 != nil || err2 != nil || limit <= 0 || offset < 0 {
		return 0, 0, &entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid query param",
			Detail:     "limit must > 0, offset >= 0 and both types are uint",
		}
	}
	return limit, offset, nil
}

func productsToBrief(products []model.Product) []ProductBrief {
	briefs := make([]ProductBrief, len(products))
	for i, p := range products {
		briefs[i] = ProductBrief{
			Id:          p.Id,
			Name:        p.Name,
			ImageUrl:    p.ImageUrl,
			Description: p.Description,
			Rating:      p.Rating,
			SoldAmount:  p.SoldAmount,
			LikedAmount: p.LikedAmount,
			CategoryId:  p.CategoryId,
			ShopId:      p.ShopId,
		}
	}
	return briefs
}

func sendProductList(ctx *gin.Context, title string, products []model.Product, limit, offset int) {
	ctx.JSON(http.StatusOK, entity.SuccessReponse[ProductListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data: ProductListResponse{
			Title:  title,
			Limit:  limit,
			Offset: offset,
			Data:   productsToBrief(products),
		},
	})
}

func GetProducts(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	limit, offset, errResp := parseLimitOffset(ctx)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	var products []model.Product
	if err := database.DB.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	sendProductList(ctx, "general", products, limit, offset)
}

func GetTopRatedProducts(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	limit, offset, errResp := parseLimitOffset(ctx)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	var products []model.Product
	if err := database.DB.Order("rating DESC").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	sendProductList(ctx, "top-rated", products, limit, offset)
}

func GetBestSellerProducts(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	limit, offset, errResp := parseLimitOffset(ctx)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	var products []model.Product
	if err := database.DB.Order("sold_amount DESC").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	sendProductList(ctx, "best-seller", products, limit, offset)
}

func GetMostLikedProducts(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	limit, offset, errResp := parseLimitOffset(ctx)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}
	var products []model.Product
	if err := database.DB.Order("liked_amount DESC").Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	sendProductList(ctx, "most-liked", products, limit, offset)
}

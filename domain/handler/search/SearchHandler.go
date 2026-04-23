package search

import (
	"WeFashionServer/domain/entity"
	"WeFashionServer/domain/handler/authentication"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// -------------------------helpers-------------------------
func ValidateTokenOrAbort(ctx *gin.Context) bool {
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

// -------------------------handler funcs-------------------------

func getQueryParams(ctx *gin.Context) (string, int, bool) {
	limitStr := ctx.DefaultQuery("limit", "10")
	query := ctx.DefaultQuery("query", "")
	limit, castErr := strconv.Atoi(limitStr)
	if castErr != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error: "Invalid param",
			Detail: "Limit must be integer"
		})
		return "", 0, false
	} 
	return query, limit, true
}

func productsToResponse(prods *[]model.Product) *[]SearchedProduct {
	p := make([]SearchedProduct, 0, len(*prods))
	for _, pr := range *prods {
		temp := SearchedProduct{
			Id:          pr.Id,
			Name:        pr.Name,
			ImageUrl:    pr.ImageUrl,
			Description: pr.Description,
			Rating:      pr.Rating,
			SoldAmount:  pr.SoldAmount,
			LikedAmount: pr.LikedAmount,
			CategoryId:  pr.CategoryId,
			ShopId:      pr.ShopId,
		}
		p = append(p, temp)
	}
	return &p
}

func shopsToResponse(shops *[]model.Shop) *[]SearchedShop {
	s := make([]SearchedShop, 0, len(*shops))
	for _, shop := range *prods {
		temp := SearchedShop{
			Id:          shop.Id,
			Name:        shop.Name,
			ImageUrl:    shop.AvatarUrl,
			Description: shop.Email,
			Rating:      shop.PhoneNumber,
			SoldAmount:  shop.Bio,
			LikedAmount: shop.RateAmount,
			CategoryId:  shop.Rating,
			ShopId:      shop.Followers,
		}
		s = append(s, temp)
	}
	return &s
}

func queryProducts(query string, limit int) *[]SearchedProduct {
	products := []model.Product{}
	if query == "" {
		database.DB.Limit(limit).Find(&products)
	} else {
		database.DB.Where("name LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Limit(limit).Find(&products)
	}
	return productsToResponse(&products)
}

func queryShops(query string, limit int) *[]SearchedShop {
	shops := []model.Shop{}
	if query == "" {
		database.DB.Limit(limit).Find(&shops)
	} else {
		database.DB.Where("name LIKE ? OR bio LIKE ?", "%"+query+"%", "%"+query+"%").Limit(limit).Find(&shops)
	}
	return shopsToResponse(&shops)
}

func SearchProductAndShop(ctx *gin.Context) {
	if !ValidateTokenOrAbort(ctx) {
		return
	}
	query, limit, success := getQueryParams(ctx)
	if !success {
		return
	}
}

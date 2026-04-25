package search

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

// -------------------------handler funcs-------------------------
func getQueryParams(ctx *gin.Context) (string, int, bool) {
	limitStr := ctx.DefaultQuery("limit", "10")
	query := ctx.DefaultQuery("query", "")
	limit, castErr := strconv.Atoi(limitStr)
	if castErr != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid param",
			Detail:     "Limit must be integer",
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
	for _, shop := range *shops {
		temp := SearchedShop{
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
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	query, limit, success := getQueryParams(ctx)
	if !success {
		return
	}
	prods := queryProducts(query, limit)
	shops := queryShops(query, limit)
	ctx.JSON(http.StatusOK, entity.SuccessReponse[interface{}]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data: gin.H{
			"products": prods,
			"shops":    shops,
		},
	})
}

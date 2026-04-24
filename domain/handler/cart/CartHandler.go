package cart

import (
	"WeFashionServer/domain/entity"
	"WeFashionServer/domain/helper"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// database.DB.Order("liked_amount DESC").Limit(limit).Offset(offset).Find(&products)

func groupSkuByProductId(items *[]CartProductSku) *map[int][]CartProductSku {
	result := make(map[int][]CartProductSku)
	for _, item := range *items {
		result[item.ProductId] = append(result[item.ProductId], item)
	}
	return &result
}

func getUser(ctx *gin.Context, userId int) *CartUser {
	user := model.User{}
	if database.DB.Where("user_id = ?", userId).Find(&user).Error != nil {
		helper.ResponseDataNotFound(
			ctx, "User", "Id", string(strconv.Itoa(123)),
		)
		return nil
	}
	return &CartUser{
		Id:          user.Id,
		Name:        user.Name,
		AvatarUrl:   user.AvatarUrl,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Bio:         user.Bio,
	}
}

func getCategory(ctx *gin.Context, id int) *CartCategory {
	category := model.Category{}
	if database.DB.Where("id = ?", id).Find(&category).Error != nil {
		helper.ResponseDataNotFound(
			ctx, "User", "Id", string(strconv.Itoa(123)),
		)
		return nil
	}
	return &CartCategory{
		Id:   category.Id,
		Name: category.Name,
	}
}

func getCartRawItem(id int) *[]CartRawItem {
	raws := []model.Cart{}
	if database.DB.Where("id = ?", id).Find(&raws).Error != nil {
		return nil
	}
	items := []CartRawItem{}
	for _, item := range raws {
		items = append(items, CartRawItem{
			SkuId:  item.SkuId,
			Amount: item.Amount,
		})
	}
	return &items
}

// Important
func GetCartProductSkus(ids *[]int) *[]CartProductSku {

	return nil
}

func GetUserCart(ctx *gin.Context) {
	if !helper.ValidateTokenOrAbort(ctx) {
		return
	}
	userId, exist := helper.GetParam[int](ctx, "id")
	if !exist {
		return
	}

	// user
	user := getUser(ctx, *userId)
	if user == nil {
		helper.ResponseRequestFieldNotFound(ctx, "Id")
	}

	// raw cart items
	raws := getCartRawItem(*userId)
	if raws == nil || len((*raws)) <= 0 {
		ctx.JSON(http.StatusOK, entity.SuccessReponse[*interface{}]{
			StatusCode: http.StatusOK,
			Time:       time.Now(),
			Data:       nil,
		})
		return
	}

	// transfer to list raw ids

}

func AddToCart(ctx *gin.Context) {

}

func RemoveFromCart(ctx *gin.Context) {

}

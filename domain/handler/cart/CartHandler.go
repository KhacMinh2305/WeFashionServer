package cart

import (
	"WeFashionServer/domain/handler/authentication"
	"WeFashionServer/domain/helper"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func groupSkuByProductId(items *[]CartProductSku) *map[int][]CartProductSku {
	result := make(map[int][]CartProductSku)
	for _, item := range *items {
		result[item.ProductId] = append(result[item.ProductId], item)
	}
	return &result
}

func getUser(ctx *gin.Context, userId int) *CartUser {
	user := model.User{}
	if database.DB.Where("id = ?", userId).Find(&user).Error != nil {
		helper.ResponseDataNotFound(
			ctx, "User", "Id", string(strconv.Itoa(userId)),
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

func getCartRawItem(id int) []CartRawItem {
	raws := []model.Cart{}
	if database.DB.Where("user_id = ?", id).Find(&raws).Error != nil {
		return nil
	}
	items := []CartRawItem{}
	for _, item := range raws {
		items = append(items, CartRawItem{
			SkuId:  item.SkuId,
			Amount: item.Amount,
		})
	}
	return items
}

func getCartProductSkus(raw []CartRawItem) ([]CartProductSku, error) {
	skuList := helper.Map(
		raw,
		func(item *CartRawItem) int {
			return item.SkuId
		},
	)
	var result []CartProductSku
	err := database.DB.Table("product_variants AS pv").
		Select(`
			pv.sku,
			pv.amount,
			pv.price,
			pv.product_id,

			s.id   AS size__id,
			s.name AS size__name,

			c.id  AS color__id,
			c.rgb AS color__rgb
		`).
		Joins("JOIN sizes s ON s.id = pv.size_id").
		Joins("JOIN colors c ON c.id = pv.color_id").
		Where("pv.sku IN ?", skuList).
		Scan(&result).Error
	return result, err
}

func querySkusByProductId(raw []CartRawItem) (map[int][]CartProductSku, error) {
	skus, err := getCartProductSkus(raw)
	if err != nil {
		return nil, err
	}
	if len(skus) != len(raw) {
		return nil, errors.New("data mismatch between raw and db result")
	}
	helper.SortBy(skus, true, func(i *CartProductSku) int {
		return i.Sku
	})
	helper.SortBy(raw, true, func(i *CartRawItem) int {
		return i.SkuId
	})
	for idx := range skus {
		if skus[idx].Sku != raw[idx].SkuId {
			return nil, errors.New("sku mismatch after sorting")
		}
		skus[idx].Amount = raw[idx].Amount
	}
	return helper.SliceToMapDuplicated(skus, func(c *CartProductSku) int {
		return c.ProductId
	}), nil
}

func GetUserCart(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	userId, exist := helper.GetParam[int](ctx, "id")
	if !exist {
		return
	}

	// user
	user := getUser(ctx, *userId)
	if user == nil {
		return
	}

	// raw cart items
	raws := getCartRawItem(*userId)
	if len(raws) == 0 {
		helper.ResponseSuccessResponse[*any](ctx, nil)
		return
	}

	skuMap, skuErr := querySkusByProductId(raws)

	if skuErr != nil {
		helper.ReponseErrorResponse(ctx, 404, "Not found", skuErr.Error())
		return
	}

	// Lấy danh sách productId từ skuMap
	productIds := make([]int, 0, len(skuMap))
	for productId := range skuMap {
		productIds = append(productIds, productId)
	}

	var products []CartProduct
	err := database.DB.Table("products AS p").
		Select(`
			p.id,
			p.name,
			p.image_url,
			p.description,
			p.rating,
			p.sold_amount,
			p.liked_amount,
			c.id AS category__id,
			c.name AS category__name
		`).
		Joins("JOIN categories c ON c.id = p.category_id").
		Where("p.id IN ?", productIds).
		Scan(&products).Error
	if err != nil {
		helper.ReponseErrorResponse(ctx, 404, "Not found", err.Error())
		return
	}

	// Map productId -> *CartProduct
	productMap := make(map[int]*CartProduct)
	for i := range products {
		productMap[products[i].Id] = &products[i]
	}
	for productId, skus := range skuMap {
		if p, ok := productMap[productId]; ok {
			p.Skus = skus
		}
	}

	helper.ResponseSuccessResponse(ctx, &CartResponse{
		User:     *user,
		Products: products,
	})
}

func validateUpdateSkuRequest(req *CartRequestUpadte) (bool, string) {
	if req.Sku <= 0 {
		return false, "sku must be a positive integer"
	}
	if req.Change != "plus" && req.Change != "minus" {
		return false, "change must be 'plus' or 'minus'"
	}
	if req.Amount <= 0 {
		return false, "amount must be a positive integer"
	}
	return true, ""
}

func getUserById(id int) (*model.User, error) {
	var user model.User
	err := database.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func getProductVariantBySku(sku int) (*model.ProductVariant, error) {
	var pv model.ProductVariant
	err := database.DB.Where("sku = ?", sku).First(&pv).Error
	if err != nil {
		return nil, err
	}
	return &pv, nil
}

func getCartItem(userId, sku int) (*model.Cart, error) {
	var cart model.Cart
	err := database.DB.Where("user_id = ? AND sku_id = ?", userId, sku).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func updateCartAmount(userId, sku, newAmount int) error {
	return database.DB.Model(&model.Cart{}).
		Where("user_id = ? AND sku_id = ?", userId, sku).
		Update("amount", newAmount).Error
}

func deleteCartItem(userId, sku int) error {
	return database.DB.Where("user_id = ? AND sku_id = ?", userId, sku).Delete(&model.Cart{}).Error
}

func UpdateSku(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}

	req := CartRequestUpadte{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ReponseErrorResponse(ctx, 400, "Invalid body", err.Error())
		return
	}
	if ok, msg := validateUpdateSkuRequest(&req); !ok {
		helper.ReponseErrorResponse(ctx, 400, "Invalid body", msg)
		return
	}

	userId, exist := helper.GetParam[int](ctx, "id")
	if !exist {
		return
	}

	user, err := getUserById(*userId)
	if err != nil {
		helper.ReponseErrorResponse(ctx, 404, "User not found", err.Error())
		return
	}
	pv, err := getProductVariantBySku(req.Sku)
	if err != nil {
		helper.ReponseErrorResponse(ctx, 404, "Sku not found", err.Error())
		return
	}
	cartItem, cartErr := getCartItem(user.Id, req.Sku)
	currentAmount := 0
	if cartErr == nil {
		currentAmount = cartItem.Amount
	}
	applied := 0
	expected := req.Amount
	if req.Change == "plus" {
		available := pv.Amount - currentAmount
		if available <= 0 {
			applied = 0
		} else if req.Amount <= available {
			applied = req.Amount
		} else {
			applied = available
		}
		if applied > 0 {
			if cartErr == nil {
				err := updateCartAmount(user.Id, req.Sku, currentAmount+applied)
				if err != nil {
					helper.ReponseErrorResponse(ctx, 500, "Update cart failed", err.Error())
					return
				}
			} else {
				err := database.DB.Create(&model.Cart{UserId: user.Id, SkuId: req.Sku, Amount: applied}).Error
				if err != nil {
					helper.ReponseErrorResponse(ctx, 500, "Create cart failed", err.Error())
					return
				}
			}
		}
	} else if req.Change == "minus" {
		if cartErr != nil || currentAmount == 0 {
			applied = 0
		} else if req.Amount < currentAmount {
			applied = req.Amount
			err := updateCartAmount(user.Id, req.Sku, currentAmount-applied)
			if err != nil {
				helper.ReponseErrorResponse(ctx, 500, "Update cart failed", err.Error())
				return
			}
		} else {
			applied = currentAmount
			err := deleteCartItem(user.Id, req.Sku)
			if err != nil {
				helper.ReponseErrorResponse(ctx, 500, "Delete cart failed", err.Error())
				return
			}
		}
	}
	resp := map[string]int{"applied": applied, "expected": expected}
	helper.ResponseSuccessResponse(ctx, &resp)
}

func RemoveFromCart(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	userId, exist := helper.GetParam[int](ctx, "id")
	if !exist {
		return
	}
	sku, exist := helper.GetQuery[int](ctx, "sku")
	if !exist {
		return
	}
	user, err := getUserById(*userId)
	if err != nil {
		helper.ReponseErrorResponse(ctx, 404, "User not found", err.Error())
		return
	}

	_, err = getCartItem(user.Id, *sku)
	if err != nil {
		helper.ResponseDataNotFound(ctx, "Products", "sku", strconv.Itoa(*sku))
		return
	}
	err = deleteCartItem(user.Id, *sku)
	if err != nil {
		helper.ReponseErrorResponse(ctx, http.StatusConflict, "Delete failed", err.Error())
		return
	}
	resp := map[string]interface{}{"success": true, "message": "Item removed from cart successfully"}
	helper.ResponseSuccessResponse(ctx, &resp)
}

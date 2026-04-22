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
	"gorm.io/gorm"
)

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
	if !ValidateTokenOrAbort(ctx) {
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
	if !ValidateTokenOrAbort(ctx) {
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
	if !ValidateTokenOrAbort(ctx) {
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
	if !ValidateTokenOrAbort(ctx) {
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

func GetProductDetail(ctx *gin.Context) {
	if !ValidateTokenOrAbort(ctx) {
		return
	}

	id, err := parseProductId(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid product id",
			Detail:     "id must be a positive integer",
		})
		return
	}

	prod, err := queryProductById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, entity.SuccessReponse[any]{
				StatusCode: http.StatusOK,
				Time:       time.Now(),
				Data:       nil,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}

	variants, err := queryProductVariants(prod.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}

	shop, err := queryShopById(prod.ShopId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}

	resp, err := buildProductDetailResponse(prod, variants, shop)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, entity.SuccessReponse[ProductDetailResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       resp,
	})
}

// --- Internal helpers for GetProductDetail ---
func parseProductId(ctx *gin.Context) (int, error) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}

func queryProductById(id int) (model.Product, error) {
	var prod model.Product
	err := database.DB.Where("id = ?", id).First(&prod).Error
	return prod, err
}

func queryProductVariants(productId int) ([]model.ProductVariant, error) {
	var variants []model.ProductVariant
	err := database.DB.Where("product_id = ?", productId).Find(&variants).Error
	return variants, err
}

func queryShopById(shopId int) (model.Shop, error) {
	var shop model.Shop
	err := database.DB.Where("id = ?", shopId).First(&shop).Error
	return shop, err
}

func buildProductDetailResponse(prod model.Product, variants []model.ProductVariant, shop model.Shop) (ProductDetailResponse, error) {
	resp := ProductDetailResponse{
		Product: ProductDetail{
			Id:          prod.Id,
			Name:        prod.Name,
			ImageUrl:    prod.ImageUrl,
			Description: prod.Description,
			Rating:      prod.Rating,
			SoldAmount:  prod.SoldAmount,
			LikedAmount: prod.LikedAmount,
			CategoryId:  prod.CategoryId,
		},
		Sku: make([]ProductVariantDetail, 0, len(variants)),
		Shop: ShopBrief{
			Id:          shop.Id,
			Name:        shop.Name,
			AvatarUrl:   shop.AvatarUrl,
			Email:       shop.Email,
			PhoneNumber: shop.PhoneNumber,
			Bio:         shop.Bio,
			RateAmount:  shop.RateAmount,
			Rating:      shop.Rating,
			Followers:   shop.Followers,
		},
	}
	for _, v := range variants {
		size, err := querySizeById(v.SizeId)
		if err != nil {
			return resp, err
		}
		color, err := queryColorById(v.ColorId)
		if err != nil {
			return resp, err
		}
		resp.Sku = append(resp.Sku, ProductVariantDetail{
			Sku:    v.Sku,
			Amount: v.Amount,
			Price:  v.Price,
			Size: SizeBrief{
				Id:   size.Id,
				Name: size.Name,
			},
			Color: ColorBrief{
				Id:  color.Id,
				Rgb: color.Rgb,
			},
		})
	}
	return resp, nil
}

func querySizeById(sizeId int) (model.Size, error) {
	var size model.Size
	err := database.DB.Where("id = ?", sizeId).First(&size).Error
	return size, err
}

func queryColorById(colorId int) (model.Color, error) {
	var color model.Color
	err := database.DB.Where("id = ?", colorId).First(&color).Error
	return color, err
}

func GetShopProducts(ctx *gin.Context) {
	if !ValidateTokenOrAbort(ctx) {
		return
	}

	shopId, err := parseShopIdForShopProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid shop id",
			Detail:     "id must be a positive integer",
		})
		return
	}

	limit, offset, errResp := parseLimitOffset(ctx)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	shop, err := queryShopByIdForShopProducts(shopId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, entity.SuccessReponse[any]{
				StatusCode: http.StatusOK,
				Time:       time.Now(),
				Data:       nil,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}

	products, err := queryProductsByShopIdForShopProducts(shopId, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}

	resp := buildShopProductsResponseForShopProducts(shop, products)
	ctx.JSON(http.StatusOK, entity.SuccessReponse[ShopProductsResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       resp,
	})
}

// --- Helpers for GetShopProducts ---
func parseShopIdForShopProducts(ctx *gin.Context) (int, error) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, err
	}
	return id, nil
}

func queryShopByIdForShopProducts(shopId int) (model.Shop, error) {
	var shop model.Shop
	err := database.DB.Where("id = ?", shopId).First(&shop).Error
	return shop, err
}

func queryProductsByShopIdForShopProducts(shopId, limit, offset int) ([]model.Product, error) {
	var products []model.Product
	err := database.DB.Where("shop_id = ?", shopId).Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

type ShopProductsResponse struct {
	Shop     ShopBrief         `json:"shop"`
	Products []ShopProductItem `json:"products"`
}

type ShopProductItem struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	ImageUrl    string  `json:"image_url"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	SoldAmount  int     `json:"sold_amount"`
	LikedAmount int     `json:"liked_amount"`
	CategoryId  int     `json:"category_id"`
}

func buildShopProductsResponseForShopProducts(shop model.Shop, products []model.Product) ShopProductsResponse {
	items := make([]ShopProductItem, len(products))
	for i, p := range products {
		items[i] = ShopProductItem{
			Id:          p.Id,
			Name:        p.Name,
			ImageUrl:    p.ImageUrl,
			Description: p.Description,
			Rating:      p.Rating,
			SoldAmount:  p.SoldAmount,
			LikedAmount: p.LikedAmount,
			CategoryId:  p.CategoryId,
		}
	}
	return ShopProductsResponse{
		Shop: ShopBrief{
			Id:          shop.Id,
			Name:        shop.Name,
			AvatarUrl:   shop.AvatarUrl,
			Email:       shop.Email,
			PhoneNumber: shop.PhoneNumber,
			Bio:         shop.Bio,
			RateAmount:  shop.RateAmount,
			Rating:      shop.Rating,
			Followers:   shop.Followers,
		},
		Products: items,
	}
}

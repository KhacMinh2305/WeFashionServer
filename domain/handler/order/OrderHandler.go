package order

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/handler/authentication"
	"WeFashionServer/domain/helper"
	"WeFashionServer/domain/repository"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"WeFashionServer/utils"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/payOSHQ/payos-lib-golang/v2"
	"gorm.io/gorm"
)

func getUserById(ctx *gin.Context, userId int) *model.User {
	user := model.User{}
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		helper.ResponseDataNotFound(ctx, "User", "id", strconv.Itoa(userId))
		return nil
	}
	return &user
}

func getOrdersByUserId(userId int) ([]orderWithProductAmountRaw, error) {
	var orders []orderWithProductAmountRaw
	err := database.DB.Table("orders AS o").
		Select(`
			o.id,
			o.discount,
			o.shipping_fee,
			o.total_price,
			o.order_state,
			o.shipping_state,
			o.created_at,
			o.user_id,
			o.address_id,
			o.payment_id,
			o.shipper_id,
			COALESCE(SUM(opv.amount), 0) AS product_amount
		`).
		Joins("LEFT JOIN order_product_variants opv ON opv.order_id = o.id").
		Where("o.user_id = ?", userId).
		Group(`
			o.id,
			o.discount,
			o.shipping_fee,
			o.total_price,
			o.order_state,
			o.shipping_state,
			o.created_at,
			o.user_id,
			o.address_id,
			o.payment_id,
			o.shipper_id
		`).
		Order("o.created_at DESC").
		Scan(&orders).Error
	return orders, err
}

func mapOrderResponse(raws []orderWithProductAmountRaw) []OrderWithProductAmountResponse {
	return helper.Map(raws, func(item *orderWithProductAmountRaw) OrderWithProductAmountResponse {
		return OrderWithProductAmountResponse{
			Id:            item.Id,
			Discount:      item.Discount,
			ShippingFee:   item.ShippingFee,
			TotalPrice:    item.TotalPrice,
			OrderState:    item.OrderState,
			ShippingState: item.ShippingState,
			CreatedAt:     item.CreatedAt,
			UserId:        item.UserId,
			AddressId:     item.AddressId,
			PaymentId:     item.PaymentId,
			ShipperId:     item.ShipperId,
			ProductAmount: item.ProductAmount,
		}
	})
}

func GetUserOrders(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}

	userId, exist := helper.GetParam[int](ctx, "id")
	if !exist {
		return
	}

	if getUserById(ctx, *userId) == nil {
		return
	}

	ordersRaw, err := getOrdersByUserId(*userId)
	if err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
		return
	}

	response := OrderListByUserResponse{
		Orders: mapOrderResponse(ordersRaw),
	}
	helper.ResponseSuccessResponse(ctx, &response)
}

func getOrderById(orderId int) (*model.Order, error) {
	order := model.Order{}
	err := database.DB.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}

func getAddressById(ctx *gin.Context, addressId int) *model.Address {
	address := model.Address{}
	if err := database.DB.Where("id = ?", addressId).First(&address).Error; err != nil {
		helper.ResponseDataNotFound(ctx, "Address", "id", strconv.Itoa(addressId))
		return nil
	}
	return &address
}

func getPaymentById(ctx *gin.Context, paymentId int) *model.Payment {
	payment := model.Payment{}
	if err := database.DB.Where("id = ?", paymentId).First(&payment).Error; err != nil {
		helper.ResponseDataNotFound(ctx, "Payment", "id", strconv.Itoa(paymentId))
		return nil
	}
	return &payment
}

func getShipperById(ctx *gin.Context, shipperId int) *model.Shipper {
	shipper := model.Shipper{}
	if err := database.DB.Where("id = ?", shipperId).First(&shipper).Error; err != nil {
		helper.ResponseDataNotFound(ctx, "Shipper", "id", strconv.Itoa(shipperId))
		return nil
	}
	return &shipper
}

func getShippingStatesByOrderId(orderId int) ([]model.ShippingState, error) {
	shippingStates := []model.ShippingState{}
	err := database.DB.Where("order_id = ?", orderId).Order("created_at ASC").Find(&shippingStates).Error
	return shippingStates, err
}

func getOrderProductsByOrderId(ctx *gin.Context, orderId int) ([]OrderDetailProductResponse, bool) {
	orderItems := []model.OrderProductVariant{}
	if err := database.DB.Where("order_id = ?", orderId).Find(&orderItems).Error; err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
		return nil, false
	}
	if len(orderItems) == 0 {
		return []OrderDetailProductResponse{}, true
	}

	skuSet := make(map[int]struct{})
	for _, item := range orderItems {
		skuSet[item.Sku] = struct{}{}
	}
	skuList := make([]int, 0, len(skuSet))
	for sku := range skuSet {
		skuList = append(skuList, sku)
	}

	variants := []model.ProductVariant{}
	if err := database.DB.Select("sku", "product_id", "size_id", "color_id").Where("sku IN ?", skuList).Find(&variants).Error; err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
		return nil, false
	}
	skuToProduct := make(map[int]int, len(variants))
	skuToSize := make(map[int]int, len(variants))
	skuToColor := make(map[int]int, len(variants))
	sizeSet := make(map[int]struct{})
	colorSet := make(map[int]struct{})
	for _, variant := range variants {
		skuToProduct[variant.Sku] = variant.ProductId
		skuToSize[variant.Sku] = variant.SizeId
		skuToColor[variant.Sku] = variant.ColorId
		sizeSet[variant.SizeId] = struct{}{}
		colorSet[variant.ColorId] = struct{}{}
	}
	sizeIds := make([]int, 0, len(sizeSet))
	for id := range sizeSet {
		sizeIds = append(sizeIds, id)
	}
	colorIds := make([]int, 0, len(colorSet))
	for id := range colorSet {
		colorIds = append(colorIds, id)
	}

	sizes := []model.Size{}
	if len(sizeIds) > 0 {
		if err := database.DB.Select("id", "name").Where("id IN ?", sizeIds).Find(&sizes).Error; err != nil {
			helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
			return nil, false
		}
	}
	sizeNames := make(map[int]string, len(sizes))
	for _, size := range sizes {
		sizeNames[size.Id] = size.Name
	}

	colors := []model.Color{}
	if len(colorIds) > 0 {
		if err := database.DB.Select("id", "rgb").Where("id IN ?", colorIds).Find(&colors).Error; err != nil {
			helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
			return nil, false
		}
	}
	colorRgbs := make(map[int]string, len(colors))
	for _, color := range colors {
		colorRgbs[color.Id] = color.Rgb
	}
	productSkus := make(map[int][]OrderDetailProductSkuResponse)
	productSet := make(map[int]struct{})
	for _, item := range orderItems {
		productId, ok := skuToProduct[item.Sku]
		if !ok {
			continue
		}
		sizeId := skuToSize[item.Sku]
		colorId := skuToColor[item.Sku]
		productSet[productId] = struct{}{}
		productSkus[productId] = append(productSkus[productId], OrderDetailProductSkuResponse{
			Sku:    item.Sku,
			Amount: item.Amount,
			Price:  item.Price,
			Size: OrderDetailSizeResponse{
				Id:   sizeId,
				Name: sizeNames[sizeId],
			},
			Color: OrderDetailColorResponse{
				Id:  colorId,
				Rgb: colorRgbs[colorId],
			},
		})
	}
	productIds := make([]int, 0, len(productSet))
	for id := range productSet {
		productIds = append(productIds, id)
	}

	products := []model.Product{}
	if err := database.DB.Where("id IN ?", productIds).Find(&products).Error; err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
		return nil, false
	}

	categorySet := make(map[int]struct{})
	shopSet := make(map[int]struct{})
	for _, product := range products {
		categorySet[product.CategoryId] = struct{}{}
		shopSet[product.ShopId] = struct{}{}
	}
	categoryIds := make([]int, 0, len(categorySet))
	for id := range categorySet {
		categoryIds = append(categoryIds, id)
	}
	shopIds := make([]int, 0, len(shopSet))
	for id := range shopSet {
		shopIds = append(shopIds, id)
	}

	categories := []model.Category{}
	if len(categoryIds) > 0 {
		if err := database.DB.Where("id IN ?", categoryIds).Find(&categories).Error; err != nil {
			helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
			return nil, false
		}
	}
	categoryNames := make(map[int]string, len(categories))
	for _, category := range categories {
		categoryNames[category.Id] = category.Name
	}

	shops := []model.Shop{}
	if len(shopIds) > 0 {
		if err := database.DB.Select("id", "name").Where("id IN ?", shopIds).Find(&shops).Error; err != nil {
			helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
			return nil, false
		}
	}
	shopNames := make(map[int]string, len(shops))
	for _, shop := range shops {
		shopNames[shop.Id] = shop.Name
	}

	responses := make([]OrderDetailProductResponse, 0, len(products))
	for _, product := range products {
		responses = append(responses, OrderDetailProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			ImageUrl:    product.ImageUrl,
			Description: product.Description,
			Rating:      product.Rating,
			SoldAmount:  product.SoldAmount,
			LikedAmount: product.LikedAmount,
			Category: OrderDetailCategoryResponse{
				Id:   product.CategoryId,
				Name: categoryNames[product.CategoryId],
			},
			Shop: OrderDetailShopResponse{
				Id:   product.ShopId,
				Name: shopNames[product.ShopId],
			},
			Skus: productSkus[product.Id],
		})
	}
	return responses, true
}

func buildOrderDetailsResponse(
	order *model.Order,
	user *model.User,
	address *model.Address,
	payment *model.Payment,
	shipper *model.Shipper,
	shippingStates []model.ShippingState,
	products []OrderDetailProductResponse,
) *OrderDetailResponse {
	return &OrderDetailResponse{
		Order: OrderDetailOrderResponse{
			Id:          order.Id,
			Discount:    order.Discount,
			ShippingFee: order.ShippingFee,
			TotalPrice:  order.TotalPrice,
			OrderState:  order.OrderState,
			CreatedAt:   order.CreatedAt,
		},
		User: OrderUserResponse{
			Id:          user.Id,
			Name:        user.Name,
			AvatarUrl:   user.AvatarUrl,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Bio:         user.Bio,
		},
		ShippingStates: helper.Map(shippingStates, func(state *model.ShippingState) OrderShippingStateResponse {
			return OrderShippingStateResponse{
				Id:        state.Id,
				Detail:    state.Detail,
				Location:  state.Location,
				Latitude:  state.Latitude,
				Longitude: state.Longitude,
				CreatedAt: state.CreatedAt,
				OrderId:   state.OrderId,
			}
		}),
		Address: OrderDetailAddressResponse{
			Id:           address.Id,
			Name:         address.Name,
			Ward:         address.Ward,
			District:     address.District,
			City:         address.City,
			Detail:       address.Detail,
			Latitude:     address.Latitude,
			Longitude:    address.Longitude,
			ReceiverName: address.ReceiverName,
			Phone:        address.Phone,
			IsDefault:    address.IsDefault,
		},
		Payment: OrderDetailPaymentResponse{
			Id:          payment.Id,
			CreatedAt:   payment.CreatedAt,
			Description: payment.Description,
			OrderId:     payment.OrderId,
		},
		Shipper: OrderShipperResponse{
			Id:          shipper.Id,
			Name:        shipper.Name,
			ImageUrl:    shipper.ImageUrl,
			PhoneNumber: shipper.PhoneNumber,
		},
		Products: products,
	}
}

func GetOrderDetails(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}

	orderId, exist := helper.GetParam[int](ctx, "id")
	if !exist {
		return
	}

	order, err := getOrderById(*orderId)
	if err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
		return
	}

	if order == nil {
		helper.ResponseSuccessResponse[*OrderDetailResponse](ctx, nil)
		return
	}

	user := getUserById(ctx, order.UserId)
	if user == nil {
		return
	}

	address := getAddressById(ctx, order.AddressId)
	if address == nil {
		return
	}

	payment := getPaymentById(ctx, order.PaymentId)
	if payment == nil {
		return
	}

	shipper := getShipperById(ctx, order.ShipperId)
	if shipper == nil {
		return
	}

	shippingStates, stateErr := getShippingStatesByOrderId(order.Id)
	if stateErr != nil {
		helper.ReponseErrorResponse(ctx, 500, "Database error", stateErr.Error())
		return
	}
	products, ok := getOrderProductsByOrderId(ctx, order.Id)
	if !ok {
		helper.ReponseErrorResponse(ctx, 500, "Database error", "")
		return
	}

	response := buildOrderDetailsResponse(order, user, address, payment, shipper, shippingStates, products)
	helper.ResponseSuccessResponse(ctx, response)
}

func HandlePaymentWebhook(ctx *gin.Context) {
	fmt.Println("----------------------------------------------------Receiver webhook----------------------------------------------------")
	webhookData := map[string]interface{}{}
	if err := ctx.ShouldBindJSON(&webhookData); err != nil {
		fmt.Println(err.Error())
		return
	}

	// ignore "ping" request for the first time to avoid testing webhook of SDK
	if checkTestWebhook(&webhookData) {
		fmt.Println("This is a test webhook from SDK.")
		return
	}

	verifiedData, err := di.PaymentRepo.VerifyPayment(webhookData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Verify webhook data valid")

	fmt.Println("--------------------Verified data--------------------")
	fmt.Println(verifiedData)
	fmt.Println("-----------------------------------------------")

	orderCode, description, ok := extractVerifiedPaymentData(verifiedData)
	if !ok {
		fmt.Println("Error parse verified data")
		return
	}

	order, shouldContinue := getOrderForPaymentWebhook(orderCode)
	if !shouldContinue {
		return
	}

	payment, err := createPaymentAndUpdateOrder(order, description)
	if err != nil {
		fmt.Println("Database error")
		return
	}

	// send email to user to notice user order is confirmed
	user := getUserById(ctx, order.UserId)
	if user == nil {
		fmt.Println("Get user failed")
		return
	}

	if err := utils.SendOrderPaidEmail(user.Email, orderCode, payment.CreatedAt); err != nil {
		fmt.Printf("Send email to %s failed!\n", user.Email)
		fmt.Println(err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{})

	fmt.Println("All is done !")

	fmt.Println("----------------------------------------------------Valid Webhook Done----------------------------------------------------")
}

func mapVerifiedPaymentData(verifiedData interface{}) (*repository.PaymentCofirmationData, bool) {
	verifiedMap, ok := verifiedData.(map[string]interface{})
	if !ok {
		return nil, false
	}
	verifiedDataRaw, ok := verifiedMap["data"]
	if !ok {
		return nil, false
	}
	verifiedDataBytes, err := json.Marshal(verifiedDataRaw)
	if err != nil {
		return nil, false
	}
	verifiedDataModel := repository.PaymentCofirmationData{}
	if err := json.Unmarshal(verifiedDataBytes, &verifiedDataModel); err != nil {
		return nil, false
	}
	return &verifiedDataModel, true
}

func extractVerifiedPaymentData(verifiedData interface{}) (int64, string, bool) {
	verifiedMap, ok := verifiedData.(map[string]interface{})
	if !ok {
		return 0, "", false
	}
	orderCodeVal, ok := verifiedMap["orderCode"]
	if !ok {
		return 0, "", false
	}
	orderCode, ok := parseOrderCode(orderCodeVal)
	if !ok {
		return 0, "", false
	}
	description, _ := verifiedMap["description"].(string)
	if description == "" {
		return 0, "", false
	}
	return orderCode, description, true
}

func parseOrderCode(value interface{}) (int64, bool) {
	switch v := value.(type) {
	case float64:
		return int64(v), true
	case float32:
		return int64(v), true
	case int:
		return int64(v), true
	case int64:
		return v, true
	case string:
		parsed, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, false
		}
		return parsed, true
	default:
		return 0, false
	}
}

func getOrderForPaymentWebhook(orderCode int64) (*model.Order, bool) {
	orderId, err := di.PaymentRepo.ResolveOrderIdFromOderCode(orderCode)
	if err != nil {
		return nil, false
	}
	order := model.Order{}
	if err := database.DB.Where("id = ?", orderId).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("Order id is: %d not found\n", orderId)
			return nil, false
		}
		fmt.Printf("Failed to find order with id: %d not found\n", orderId)
		return nil, false
	}
	if order.OrderState == 1 {
		fmt.Printf("Order %d were updated!\n", orderId)
		return nil, false
	}
	return &order, true
}

func createPaymentAndUpdateOrder(order *model.Order, description string) (*model.Payment, error) {
	payment := model.Payment{
		CreatedAt:   time.Now(),
		Description: description,
		OrderId:     order.Id,
		UserId:      order.UserId,
	}

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&payment).Error; err != nil {
			fmt.Println("Created payment record failed !")
			return err
		}
		fmt.Println("Created payment record successfully !")
		if err := tx.Model(order).Updates(map[string]interface{}{
			"order_state": 1,
			"payment_id":  payment.Id,
		}).Error; err != nil {
			fmt.Println("Updated order state to confirmed failed !")
			return err
		}
		fmt.Println("Updated order state to confirmed successfully !")
		return nil
	}); err != nil {
		fmt.Println("Transaction is broken when trying to change order state !")
		return nil, err
	}
	fmt.Printf("All Update for order %d successfully!\n", order.Id)
	return &payment, nil
}

func checkTestWebhook(body *map[string]interface{}) bool {
	data, ok := (*body)["data"].(map[string]interface{})
	if !ok {
		return false
	}

	orderCodeVal, ok := data["orderCode"].(float64)
	if !ok {
		return false
	}
	orderCode := int64(orderCodeVal)
	accountNumber, _ := data["accountNumber"].(string)
	paymentLinkId, _ := data["paymentLinkId"].(string)
	return orderCode == 123 && accountNumber == "12345678" && paymentLinkId == "124c33293c43417ab7879e14c8d9eb18"
}

func validateOrderItem(item *OrderItem) (bool, string) {
	if item.Sku <= 0 {
		return false, "sku must be a positive integer"
	}
	if item.Amount <= 0 {
		return false, "amount must be a positive integer"
	}
	if item.Price <= 0 {
		return false, "price must be a positive number"
	}
	if item.ProductId <= 0 {
		return false, "product_id must be a positive integer"
	}
	if item.SizeId <= 0 {
		return false, "size_id must be a positive integer"
	}
	if item.ColorId <= 0 {
		return false, "color_id must be a positive integer"
	}
	return true, ""
}

func validateOrderRequestBody(req *OrderRequestBody) (bool, string) {
	if req.Discount <= 0 {
		return false, "discount must be a positive number"
	}
	if req.ShippingFee <= 0 {
		return false, "shipping_fee must be a positive number"
	}
	if req.TotalPrice <= 0 {
		return false, "total_price must be a positive number"
	}
	if req.UserId <= 0 {
		return false, "user_id must be a positive integer"
	}
	if req.AddressId <= 0 {
		return false, "address_id must be a positive integer"
	}
	if len(req.Items) == 0 {
		return false, "items must not be empty"
	}
	for idx := range req.Items {
		if ok, msg := validateOrderItem(&req.Items[idx]); !ok {
			return false, fmt.Sprintf("items[%d]: %s", idx, msg)
		}
	}
	return true, ""
}

func getAddressByIdForUser(ctx *gin.Context, addressId, userId int) *model.Address {
	address := getAddressById(ctx, addressId)
	if address == nil {
		return nil
	}
	if address.UserId != userId {
		helper.ReponseErrorResponse(ctx, 400, "Invalid address", "address does not belong to user")
		return nil
	}
	return address
}

func validateSkusExist(items []OrderItem) (bool, error) {
	skuSet := make(map[int]struct{})
	for _, item := range items {
		skuSet[item.Sku] = struct{}{}
	}
	if len(skuSet) == 0 {
		return false, nil
	}
	skuList := make([]int, 0, len(skuSet))
	for sku := range skuSet {
		skuList = append(skuList, sku)
	}
	variants := []model.ProductVariant{}
	if err := database.DB.Select("sku").Where("sku IN ?", skuList).Find(&variants).Error; err != nil {
		return false, err
	}
	if len(variants) != len(skuList) {
		return false, nil
	}
	return true, nil
}

func validateSkusInStock(items []OrderItem) (bool, error) {
	skuSet := make(map[int]struct{})
	for _, item := range items {
		skuSet[item.Sku] = struct{}{}
	}
	if len(skuSet) == 0 {
		return false, nil
	}
	skuList := make([]int, 0, len(skuSet))
	for sku := range skuSet {
		skuList = append(skuList, sku)
	}
	variants := []model.ProductVariant{}
	if err := database.DB.Select("sku", "amount").Where("sku IN ?", skuList).Find(&variants).Error; err != nil {
		return false, err
	}
	stockBySku := make(map[int]int, len(variants))
	for _, variant := range variants {
		stockBySku[variant.Sku] = variant.Amount
	}
	for _, item := range items {
		if stockBySku[item.Sku] < item.Amount {
			return false, nil
		}
	}
	return true, nil
}

func getUniqueProductIds(items []OrderItem) []int {
	productSet := make(map[int]struct{})
	for _, item := range items {
		if item.ProductId > 0 {
			productSet[item.ProductId] = struct{}{}
		}
	}
	productIds := make([]int, 0, len(productSet))
	for id := range productSet {
		productIds = append(productIds, id)
	}
	return productIds
}

func getProductNamesByIds(ctx *gin.Context, productIds []int) (map[int]string, bool) {
	if len(productIds) == 0 {
		return map[int]string{}, true
	}
	products := []model.Product{}
	if err := database.DB.Select("id", "name").Where("id IN ?", productIds).Find(&products).Error; err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
		return nil, false
	}
	if len(products) != len(productIds) {
		helper.ReponseErrorResponse(ctx, 404, "Product not found", "one or more product not found")
		return nil, false
	}
	nameMap := make(map[int]string, len(products))
	for _, product := range products {
		nameMap[product.Id] = product.Name
	}
	return nameMap, true
}

func buildPaymentItems(ctx *gin.Context, items []OrderItem, productNames map[int]string) ([]payos.PaymentLinkItem, bool) {
	paymentItems := make([]payos.PaymentLinkItem, 0, len(items))
	for _, item := range items {
		name, exist := productNames[item.ProductId]
		if !exist {
			helper.ReponseErrorResponse(ctx, 404, "Product not found", "one or more product not found")
			return nil, false
		}
		paymentItems = append(paymentItems, payos.PaymentLinkItem{
			Name:     name,
			Quantity: item.Amount,
			Price:    int(item.Price),
		})
	}
	return paymentItems, true
}

func getRandomShipper(ctx *gin.Context) *model.Shipper {
	shippers := []model.Shipper{}
	if err := database.DB.Find(&shippers).Error; err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
		return nil
	}
	if len(shippers) == 0 {
		helper.ReponseErrorResponse(ctx, 404, "Shipper not found", "no shipper available")
		return nil
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &shippers[rng.Intn(len(shippers))]
}

func CreateOrder(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}

	userId, exist := helper.GetParam[int](ctx, "id")
	if !exist {
		return
	}
	if *userId <= 0 {
		helper.ReponseErrorResponse(ctx, 400, "Invalid param", "id must be a positive integer")
		return
	}

	user := getUserById(ctx, *userId)
	if user == nil {
		return
	}

	req := OrderRequestBody{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ReponseErrorResponse(ctx, 400, "Invalid body", err.Error())
		return
	}
	if ok, msg := validateOrderRequestBody(&req); !ok {
		helper.ReponseErrorResponse(ctx, 400, "Invalid body", msg)
		return
	}

	if req.UserId != *userId {
		helper.ReponseErrorResponse(ctx, 400, "Invalid body", "user_id must match id in param")
		return
	}

	address := getAddressByIdForUser(ctx, req.AddressId, req.UserId)
	if address == nil {
		return
	}

	if ok, err := validateSkusExist(req.Items); err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
		return
	} else if !ok {
		helper.ReponseErrorResponse(ctx, 404, "Sku not found", "one or more sku not found")
		return
	}
	if ok, err := validateSkusInStock(req.Items); err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Database error", err.Error())
		return
	} else if !ok {
		helper.ReponseErrorResponse(ctx, 400, "Sku not enough stock", "sku does not have enough stock, please check again")
		return
	}

	shipper := getRandomShipper(ctx)
	if shipper == nil {
		return
	}

	order := model.Order{
		Discount:      req.Discount,
		ShippingFee:   req.ShippingFee,
		TotalPrice:    req.TotalPrice,
		OrderState:    0,
		ShippingState: -1,
		CreatedAt:     time.Now(),
		UserId:        req.UserId,
		AddressId:     address.Id,
		PaymentId:     -1,
		ShipperId:     shipper.Id,
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		items := make([]model.OrderProductVariant, 0, len(req.Items))
		for _, item := range req.Items {
			items = append(items, model.OrderProductVariant{
				OrderId: order.Id,
				Sku:     item.Sku,
				Amount:  item.Amount,
				Price:   item.Price,
			})
		}
		if len(items) > 0 {
			if err := tx.Create(&items).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Create order failed", err.Error())
		return
	}

	productIds := getUniqueProductIds(req.Items)
	productNames, ok := getProductNamesByIds(ctx, productIds)
	if !ok {
		return
	}
	paymentItems, ok := buildPaymentItems(ctx, req.Items, productNames)
	if !ok {
		return
	}

	buyerAddress := strings.Join([]string{address.City, address.District, address.Ward, address.Detail}, "-")
	paymentData := repository.PaymentData{
		OrderCode:    int64(order.Id),
		Amount:       int(2000), /*req.TotalPrice*/
		Description:  "WeFashion Payment",
		CancelUrl:    "https://www.google.com/?hl=vi",
		ReturnUrl:    "https://go.dev/",
		Items:        paymentItems,
		BuyerName:    &user.Name,
		BuyerEmail:   &user.Email,
		BuyerPhone:   &user.PhoneNumber,
		BuyerAddress: &buyerAddress,
	}

	paymentLink, err := di.PaymentRepo.CreatePaymentRequest(paymentData)
	if err != nil {
		helper.ReponseErrorResponse(ctx, 500, "Create payment link failed", err.Error())
		return
	}

	response := OrderPaymentLinkResponse{
		CheckoutUrl: paymentLink.CheckoutUrl,
		QrCode:      paymentLink.QrCode,
	}
	helper.ResponseSuccessResponse(ctx, &response)
}

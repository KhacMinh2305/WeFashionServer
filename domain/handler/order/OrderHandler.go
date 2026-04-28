package order

import (
	"WeFashionServer/domain/handler/authentication"
	"WeFashionServer/domain/helper"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
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

func buildOrderDetailsResponse(
	order *model.Order,
	user *model.User,
	address *model.Address,
	payment *model.Payment,
	shipper *model.Shipper,
	shippingStates []model.ShippingState,
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

	response := buildOrderDetailsResponse(order, user, address, payment, shipper, shippingStates)
	helper.ResponseSuccessResponse(ctx, response)
}

func HandlePaymentWebhook(ctx *gin.Context) {
	fmt.Println("----------------------Receiver webhook----------------------")
}

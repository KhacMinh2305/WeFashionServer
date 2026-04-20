package address

import (
	"WeFashionServer/domain/entity"
	"WeFashionServer/domain/handler/authentication"
	"WeFashionServer/infrastructure/database"
	"WeFashionServer/infrastructure/model"
	"net/http"
	"strconv"
	"strings"
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

// GET /api/address/user/:id
func GetAddressesByUserId(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	idStr := ctx.Param("id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid user id",
			Detail:     "User id must be an integer.",
		})
		return
	}
	// Lấy user
	var user model.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "User not found",
			Detail:     "No user with this id.",
		})
		return
	}
	// Lấy danh sách addresses
	var addresses []model.Address
	if err := database.DB.Where("user_id = ?", userId).Find(&addresses).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[map[string]interface{}]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data: map[string]interface{}{
			"user":      user,
			"addresses": addresses,
		},
	})
}

// GET /api/address/:id
func GetAddressById(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	idStr := ctx.Param("id")
	addressId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid address id",
			Detail:     "Address id must be an integer.",
		})
		return
	}
	var address model.Address
	if err := database.DB.Where("id = ?", addressId).First(&address).Error; err != nil {
		// Không tìm thấy thì trả về success, data = null
		ctx.JSON(http.StatusOK, entity.SuccessReponse[*model.Address]{
			StatusCode: http.StatusOK,
			Time:       time.Now(),
			Data:       nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[model.Address]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       address,
	})
}

// Helper: validate các trường của address request
func validateAddressRequest(req *struct {
	Name         string  `json:"name"`
	Ward         string  `json:"ward"`
	District     string  `json:"district"`
	City         string  `json:"city"`
	Detail       string  `json:"detail"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	ReceiverName string  `json:"receiver_name"`
	Phone        string  `json:"phone"`
	IsDefault    bool    `json:"is_default"`
	UserId       int     `json:"user_id"`
}) (string, bool) {
	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Ward) == "" || strings.TrimSpace(req.District) == "" ||
		strings.TrimSpace(req.City) == "" || strings.TrimSpace(req.Detail) == "" ||
		strings.TrimSpace(req.ReceiverName) == "" || strings.TrimSpace(req.Phone) == "" || req.UserId <= 0 {
		return "All fields must be non-empty and user_id must be valid.", false
	}
	if req.Latitude < -90 || req.Latitude > 90 || req.Longitude < -180 || req.Longitude > 180 {
		return "Latitude must be in [-90,90], longitude in [-180,180]", false
	}
	return "", true
}

// Helper: kiểm tra user có tồn tại không
func userExists(userId int) bool {
	var user model.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return false
	}
	return true
}

// POST /api/address/create
func CreateAddress(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	var req struct {
		Name         string  `json:"name"`
		Ward         string  `json:"ward"`
		District     string  `json:"district"`
		City         string  `json:"city"`
		Detail       string  `json:"detail"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
		ReceiverName string  `json:"receiver_name"`
		Phone        string  `json:"phone"`
		IsDefault    bool    `json:"is_default"`
		UserId       int     `json:"user_id"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid request body",
			Detail:     err.Error(),
		})
		return
	}
	// Validate fields bằng helper
	if msg, ok := validateAddressRequest(&req); !ok {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid input",
			Detail:     msg,
		})
		return
	}
	// Kiểm tra user tồn tại bằng helper
	if !userExists(req.UserId) {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "User not found",
			Detail:     "No user with this id.",
		})
		return
	}
	// Tạo address mới
	address := model.Address{
		Name:         req.Name,
		Ward:         req.Ward,
		District:     req.District,
		City:         req.City,
		Detail:       req.Detail,
		Latitude:     req.Latitude,
		Longitude:    req.Longitude,
		ReceiverName: req.ReceiverName,
		Phone:        req.Phone,
		IsDefault:    req.IsDefault,
		UserId:       req.UserId,
	}
	if err := database.DB.Create(&address).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Create address failed",
			Detail:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[model.Address]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       address,
	})
}

// PUT /api/address/update
func UpdateAddress(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	addressIdStr := ctx.Query("address_id")
	addressId, err := strconv.Atoi(addressIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid address_id",
			Detail:     "address_id must be an integer.",
		})
		return
	}
	var req struct {
		Name         string  `json:"name"`
		Ward         string  `json:"ward"`
		District     string  `json:"district"`
		City         string  `json:"city"`
		Detail       string  `json:"detail"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
		ReceiverName string  `json:"receiver_name"`
		Phone        string  `json:"phone"`
		IsDefault    bool    `json:"is_default"`
		UserId       int     `json:"user_id"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid request body",
			Detail:     err.Error(),
		})
		return
	}
	// Validate fields
	if msg, ok := validateAddressRequest(&req); !ok {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid input",
			Detail:     msg,
		})
		return
	}
	// Check address exists
	var address model.Address
	if err := database.DB.Where("id = ?", addressId).First(&address).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Address not found",
			Detail:     "No address with this id.",
		})
		return
	}
	// Check address_id phải trùng với req.UserId
	if address.Id != addressId || address.UserId != req.UserId {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Id mismatch",
			Detail:     "conflict address id",
		})
		return
	}
	// Check user exists
	if !userExists(req.UserId) {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "User not found",
			Detail:     "No user with this id.",
		})
		return
	}
	// Update address fields
	address.Name = req.Name
	address.Ward = req.Ward
	address.District = req.District
	address.City = req.City
	address.Detail = req.Detail
	address.Latitude = req.Latitude
	address.Longitude = req.Longitude
	address.ReceiverName = req.ReceiverName
	address.Phone = req.Phone
	address.IsDefault = req.IsDefault
	address.UserId = req.UserId
	if err := database.DB.Save(&address).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Update address failed",
			Detail:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[model.Address]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       address,
	})
}

// DELETE /api/address/:id/delete
func DeleteAddress(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	idStr := ctx.Param("id")
	addressId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid address id",
			Detail:     "Address id must be an integer.",
		})
		return
	}
	var address model.Address
	if err := database.DB.Where("id = ?", addressId).First(&address).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Address not found",
			Detail:     "No address with this id.",
		})
		return
	}
	if err := database.DB.Delete(&address).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Delete address failed",
			Detail:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[map[string]interface{}]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data: map[string]interface{}{
			"message": "Address deleted successfully",
		},
	})
}

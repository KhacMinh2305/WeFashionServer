package user

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

// GET /api/user?user_id=123
func GetUser(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	userIdStr := ctx.Query("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid user_id",
			Detail:     "user_id must be an integer.",
		})
		return
	}
	var user model.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		ctx.JSON(http.StatusOK, entity.SuccessReponse[*model.User]{
			StatusCode: http.StatusOK,
			Time:       time.Now(),
			Data:       nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[model.User]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       user,
	})
}

// POST /api/user/:id/update
func UpdateUser(ctx *gin.Context) {
	if !validateTokenOrAbort(ctx) {
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid id param",
			Detail:     "id must be an integer.",
		})
		return
	}
	var req struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		AvatarUrl   string `json:"avatar_url"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Bio         string `json:"bio"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Invalid request body",
			Detail:     err.Error(),
		})
		return
	}
	if req.Id != id {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Id mismatch",
			Detail:     "id in param and body must match.",
		})
		return
	}
	var user model.User
	if err := database.DB.Where("id = ?", req.Id).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "User not found",
			Detail:     "No user with this id.",
		})
		return
	}
	user.Name = req.Name
	user.AvatarUrl = req.AvatarUrl
	user.Email = req.Email
	user.PhoneNumber = req.PhoneNumber
	user.Bio = req.Bio
	if err := database.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Update failed",
			Detail:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[model.User]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       user,
	})
}

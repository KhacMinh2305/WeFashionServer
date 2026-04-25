package color

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

// GET /api/color : get all colors or by id if color_id param exists
func GetColors(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	idStr := ctx.Query("color_id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Error:      "Invalid color_id",
				Detail:     err.Error(),
			})
			return
		}
		var colorModel model.Color
		if err := database.DB.First(&colorModel, id).Error; err != nil {
			// Not found: return success with data = null
			ctx.JSON(http.StatusOK, entity.SuccessReponse[*ColorResponse]{
				StatusCode: http.StatusOK,
				Time:       time.Now(),
				Data:       nil,
			})
			return
		}
		resp := ColorResponse{
			Id:  colorModel.Id,
			Rgb: colorModel.Rgb,
		}
		ctx.JSON(http.StatusOK, entity.SuccessReponse[*ColorResponse]{
			StatusCode: http.StatusOK,
			Time:       time.Now(),
			Data:       &resp,
		})
		return
	}
	// Nếu không có param thì trả về tất cả
	var colors []model.Color
	if err := database.DB.Find(&colors).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	var resp []ColorResponse
	for _, c := range colors {
		resp = append(resp, ColorResponse{
			Id:  c.Id,
			Rgb: c.Rgb,
		})
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[ColorListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       ColorListResponse{Colors: resp},
	})
}

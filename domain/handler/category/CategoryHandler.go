package category

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

// GET /api/category : get all categories or by id if category_id param exists
func GetCategories(ctx *gin.Context) {
	if !authentication.ValidateTokenOrAbort(ctx) {
		return
	}
	idStr := ctx.Query("category_id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Error:      "Invalid category_id",
				Detail:     err.Error(),
			})
			return
		}
		var categoryModel model.Category
		if err := database.DB.First(&categoryModel, id).Error; err != nil {
			// Not found: return success with data = null
			ctx.JSON(http.StatusOK, entity.SuccessReponse[*CategoryResponse]{
				StatusCode: http.StatusOK,
				Time:       time.Now(),
				Data:       nil,
			})
			return
		}
		resp := CategoryResponse{
			Id:   categoryModel.Id,
			Name: categoryModel.Name,
		}
		ctx.JSON(http.StatusOK, entity.SuccessReponse[*CategoryResponse]{
			StatusCode: http.StatusOK,
			Time:       time.Now(),
			Data:       &resp,
		})
		return
	}
	// Nếu không có param thì trả về tất cả
	var categories []model.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Database error",
			Detail:     err.Error(),
		})
		return
	}
	var resp []CategoryResponse
	for _, c := range categories {
		resp = append(resp, CategoryResponse{
			Id:   c.Id,
			Name: c.Name,
		})
	}
	ctx.JSON(http.StatusOK, entity.SuccessReponse[CategoryListResponse]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       CategoryListResponse{Categories: resp},
	})
}

package helper

import (
	"WeFashionServer/domain/entity"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ---------------------------code---------------------------
const UNAUTHORIZED_CODE = http.StatusUnauthorized
const NOT_FOUND_CODE = http.StatusNotFound
const PARSE_CODE = http.StatusConflict

// ---------------------------Error---------------------------
const UNAUTHORIZED_ERROR = "Invalid token"
const NOT_FOUND_ERROR = "Not found"
const PARSE_ERROR = "Invalid type"

// ---------------------------Detail---------------------------
const UNAUTHORIZED_DETAIL = "Token may be experied or invalid"
const NOT_FOUND_DETAIL_REQUEST = "%s not found"
const NOT_FOUND_DETAIL_DATA = "%s not found with the %s %s"
const PARSE_DETAIL = "%s must be %s"

func buildErrorResponse(code int, err, detail string) *entity.ErrorResponse {
	return &entity.ErrorResponse{
		StatusCode: code,
		Error:      err,
		Detail:     detail,
	}
}

func ReponseErrorResponse(ctx *gin.Context, code int, err, detail string) {
	ctx.JSON(
		code,
		buildErrorResponse(code, err, detail),
	)
}

func ResponseSuccessResponse[T *any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, entity.SuccessReponse[T]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       data,
	})
}

func ResponseUnauthorized(ctx *gin.Context, detail *string) {
	errDetail := UNAUTHORIZED_DETAIL
	if detail != nil {
		errDetail = *detail
	}
	ReponseErrorResponse(
		ctx,
		UNAUTHORIZED_CODE,
		UNAUTHORIZED_ERROR,
		errDetail,
	)
}

func ResponseRequestFieldNotFound(ctx *gin.Context, field string) {
	ReponseErrorResponse(
		ctx,
		NOT_FOUND_CODE,
		NOT_FOUND_ERROR,
		fmt.Sprintf(NOT_FOUND_DETAIL_REQUEST, field),
	)
}

func ResponseParseError(ctx *gin.Context, field, typeData string) {
	ReponseErrorResponse(
		ctx,
		PARSE_CODE,
		PARSE_ERROR,
		fmt.Sprintf(PARSE_DETAIL, field, typeData),
	)
}

func ResponseDataNotFound(ctx *gin.Context, dataType, field, value string) {
	ReponseErrorResponse(
		ctx,
		NOT_FOUND_CODE,
		NOT_FOUND_ERROR,
		fmt.Sprintf(NOT_FOUND_DETAIL_DATA, dataType, field, value),
	)
}

package helper

import (
	"WeFashionServer/domain/entity"
	"WeFashionServer/domain/handler/authentication"
	"net/http"

	"github.com/gin-gonic/gin"
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

func ParseRequestExtraInfo[T any](key string, ctx *gin.Context, keyProvider func() (string, bool)) (*T, bool) {
	valueStr, exist := keyProvider()
	if !exist {
		ResponseRequestFieldNotFound(ctx, key)
		return nil, false
	}
	value, parseErr := ParsePrimitive[T](valueStr)
	if parseErr != nil {
		ResponseParseError(ctx, key, GetTypeName[T]())
		return nil, false
	}
	return &value, true
}

func GetParam[T any](ctx *gin.Context, key string) (*T, bool) {
	return ParseRequestExtraInfo[T](
		key, ctx,
		func() (string, bool) {
			return ctx.GetParam(key)
		},
	)
}

func GetQuery[T any](ctx *gin.Context, key string) (*T, bool) {
	return ParseRequestExtraInfo[T](
		key, ctx,
		func() (string, bool) {
			return ctx.GetQuery(key)
		},
	)
}

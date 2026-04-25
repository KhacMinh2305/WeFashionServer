package helper

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func parseRequestExtraInfo[T any](key string, ctx *gin.Context, keyProvider func() (string, bool)) (*T, bool) {
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
	return parseRequestExtraInfo[T](
		key, ctx,
		func() (string, bool) {
			keyStr := strings.Trim(ctx.Param(key), "/")
			return keyStr, keyStr != ""
		},
	)
}

func GetQuery[T any](ctx *gin.Context, key string) (*T, bool) {
	return parseRequestExtraInfo[T](
		key, ctx,
		func() (string, bool) {
			return ctx.GetQuery(key)
		},
	)
}

package authentication

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/entity"
	"WeFashionServer/domain/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FetchToken(ctx *gin.Context) {
	newToken, genErr := di.AuthRepo.GenerateToken()
	if genErr != nil {
		helper.ReponseErrorResponse(ctx, http.StatusUnauthorized, "Cannot generate token", genErr.Error())
		return
	}

	ctx.JSON(http.StatusOK, entity.SuccessReponse[*entity.EntityToken]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       newToken,
	})
}

func validateToken(ctx *gin.Context) (bool, *string) {
	var err string
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		err = "Authorization header missing"
		return false, &err
	}
	var tokenString string
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	} else {
		err = "Invalid Authorization header format"
		return false, &err
	}
	parsedToken, parseErr := di.AuthRepo.ParseToken(tokenString)
	if parseErr != nil {
		err = "Cannot parse token: " + parseErr.Error()
		return false, &err
	}
	if !parsedToken.Valid {
		err = "Invalid token"
		return false, &err
	}
	return true, nil
}

func ValidateTokenOrAbort(ctx *gin.Context) bool {
	valid, errStr := validateToken(ctx)
	if !valid {
		helper.ResponseUnauthorized(ctx, errStr)
		return false
	}
	return true
}

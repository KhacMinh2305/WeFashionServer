package authentication

import (
	"WeFashionServer/di"
	"WeFashionServer/domain/entity"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenRequestBody struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

func createUnauthorizedResponse(err error, detail string) *entity.ErrorResponse {
	return &entity.ErrorResponse{
		StatusCode: http.StatusUnauthorized,
		Error:      err.Error(),
		Detail:     detail,
	}
}

func FetchToken(ctx *gin.Context) {
	body := &AccessTokenRequestBody{}

	mapErr := ctx.ShouldBindJSON(body)
	if mapErr != nil {
		ctx.JSON(http.StatusUnauthorized, createUnauthorizedResponse(mapErr, "Cannot validate token!"))
		return
	}

	if body.Token != "" {
		parsedToken, parseErr := di.AuthRepo.ParseToken(body.Token)
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); parseErr != nil || !ok {
			ctx.JSON(http.StatusUnauthorized, createUnauthorizedResponse(parseErr, "Cannot decode token!"))
			return
		} else {
			if field, ok := claims["user_id"].(string); ok {
				if !parsedToken.Valid || field != body.Id {
					detail := ""
					if field == body.Id {
						detail = "Token may be expired or incorrect!"
					} else {
						detail = "Token is not belong to " + body.Id
					}
					ctx.JSON(http.StatusUnauthorized, createUnauthorizedResponse(errors.New("Invalid token"), detail))
					return
				}
			} else {
				ctx.JSON(http.StatusUnauthorized, createUnauthorizedResponse(errors.New("Invalid token"), ""))
			}
		}
	}

	newToken, genErr := di.AuthRepo.GenerateToken(body.Id)
	if genErr != nil {
		ctx.JSON(http.StatusUnauthorized, createUnauthorizedResponse(genErr, "Cannot generate access token now!"))
		return
	}

	ctx.JSON(http.StatusOK, entity.SuccessReponse[*entity.EntityToken]{
		StatusCode: http.StatusOK,
		Time:       time.Now(),
		Data:       newToken,
	})
}

// Use to validate token in header of request, return user_id if valid, otherwise return error
func ValidateToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header missing")
	}

	var tokenString string
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	} else {
		return "", errors.New("Invalid Authorization header format")
	}

	parsedToken, parseErr := di.AuthRepo.ParseToken(tokenString)
	if parseErr != nil {
		return "", errors.New("Cannot parse token: " + parseErr.Error())
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", errors.New("Invalid token claims or token is not valid")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("user_id not found in token claims")
	}

	return userID, nil
}

package repository

import (
	"WeFashionServer/domain/entity"

	"github.com/golang-jwt/jwt/v5"
)

type AuthRepository interface {
	GenerateToken(targetId string) (*entity.EntityToken, error)
	ParseToken(token string) (*jwt.Token, error)
}

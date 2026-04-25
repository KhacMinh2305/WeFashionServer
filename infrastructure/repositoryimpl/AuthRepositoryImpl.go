package repositoryimpl

import (
	"WeFashionServer/domain/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthRepositoryImpl struct{}

func (repo *AuthRepositoryImpl) GenerateToken( /*targetId string*/ ) (*entity.EntityToken, error) {
	privateKey := []byte(os.Getenv("JWT_PRIVATE_KEY"))
	expiredAt := time.Now().Add(time.Hour * 1)
	createdAt := time.Now()
	claims := jwt.MapClaims{
		// "user_id": targetId,
		"exp": expiredAt.Unix(),
		"iat": createdAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return nil, err
	}
	return &entity.EntityToken{
		Token:     tokenString,
		CreatedAt: createdAt.UnixMilli(),
		ExpiredIn: 3600,
	}, nil
}

func (repo *AuthRepositoryImpl) ParseToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_PRIVATE_KEY")), nil
	})
	return parsedToken, err
}

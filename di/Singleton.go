package di

import (
	"WeFashionServer/domain/repository"
	"WeFashionServer/infrastructure/repositoryimpl"

	"github.com/gin-gonic/gin"
)

// route
var Router = gin.Default()

// repository
var AuthRepo repository.AuthRepository = &repositoryimpl.AuthRepositoryImpl{}

package di

import (
	"WeFashionServer/domain/repository"
	"WeFashionServer/infrastructure/repositoryimpl"

	"github.com/gin-gonic/gin"
)

// repository
var AuthRepo repository.AuthRepository = &repositoryimpl.AuthRepositoryImpl{}

// route
var Router = gin.Default()

package api

import (
	"go-gin-album/docs"
	"go-gin-album/internal/constant"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpAPIDocs(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = constant.BasePath
	router.GET(constant.Swagger, ginSwagger.WrapHandler(swaggerfiles.Handler))
}

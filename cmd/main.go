package main

import (
	"go-gin-album/api"
	"go-gin-album/internal/server"

	"github.com/gin-gonic/gin"
)

// @title           Music Album API
// @version         1.0
// @description     API documentation for Music Album service
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://joshydavid.com
// @contact.email  joshuadavidang@outlook.sg

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	api.SetupRoutes(router)
	api.SetUpAPIDocs(router)
	server.RunServer(router)
}

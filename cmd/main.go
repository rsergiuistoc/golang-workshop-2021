package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/routes"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// @title Workshop Todo-Service Api
// @version 1.0
// @description Swagger API for Workshop Todo-Service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func main(){

	cfg := internal.InitConfiguration(".env", ".")
	db := internal.CreateDatabaseConn(cfg)

	router := gin.Default()

	api := router.Group("/api")
	{
		routes.ApplyStatusRoutes(api)
		routes.ApplyAuthenticationRoutes(api, db, cfg)
		routes.ApplyTodoRoutes(api, db, cfg)
	}

	// Error Endpoints
	router.NoRoute(func(c *gin.Context){
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "Resource Not Found"})
	})

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	PORT := "5000"
	router.Run(fmt.Sprintf(":%s", PORT))
}
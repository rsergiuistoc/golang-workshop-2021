package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/routes"
	"net/http"
)

func main(){

	cfg := internal.InitConfiguration(".env", ".")
	_ = internal.CreateDatabaseConn(cfg)

	router := gin.Default()

	api := router.Group("/api")
	{
		routes.ApplyStatusRoutes(api)
	}

	// Error Endpoints
	router.NoRoute(func(c *gin.Context){
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "Resource Not Found"})
	})

	PORT := "5000"
	router.Run(fmt.Sprintf(":%s", PORT))
}
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/controllers"
)

// ApplyStatusRoutes attaches status endpoints: ping/metrics/health
//to the root group.
func ApplyStatusRoutes(r *gin.RouterGroup){

	statusApi := r.Group("/status")
	{
		statusApi.GET("/ping", controllers.Ping)
	}
}
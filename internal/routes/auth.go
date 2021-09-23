package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/controllers"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/middlewares"
	"gorm.io/gorm"
)

// ApplyAuthenticationRoutes attaches authentication endpoints to
// the root group.
func ApplyAuthenticationRoutes(r *gin.RouterGroup, d *gorm.DB, cfg *internal.Configuration) {

	controller := controllers.NewAuthController(d, cfg)

	auth := r.Group("/auth")
	{
		auth.POST("/signin", middlewares.BasicAuthentication(d), controller.SignIn)
		auth.POST("/signup", controller.SignUp)
	}
}
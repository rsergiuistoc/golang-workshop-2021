package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/controllers"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/middlewares"
	"gorm.io/gorm"
)

func ApplyTodoRoutes(r *gin.RouterGroup, d *gorm.DB, cfg *internal.Configuration){

	controller := controllers.NewTodoController(d)

	todo := r.Group("/todos")
	{
		todo.GET("/", middlewares.AuthorizeToken(cfg), controller.ListTodos)
		todo.GET("/:id", middlewares.AuthorizeToken(cfg), controller.RetrieveTodo)
		todo.POST("/", middlewares.AuthorizeToken(cfg), controller.CreateTodo)
		todo.PATCH("/:id", middlewares.AuthorizeToken(cfg), controller.UpdateTodo)
		todo.DELETE("/:id", middlewares.AuthorizeToken(cfg), controller.DeleteTodo)
	}
}
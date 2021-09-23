package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/controllers"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/middlewares"
	"gorm.io/gorm"
)

func ApplyTodoRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *internal.Configuration){

	controller := controllers.NewTodoController(db)

	todo := r.Group("/todos")
	{
		todo.GET("/", middlewares.AuthorizeToken(db, cfg), controller.ListTodos)
		todo.GET("/:id", middlewares.AuthorizeToken(db, cfg), controller.RetrieveTodo)
		todo.POST("/", middlewares.AuthorizeToken(db, cfg), controller.CreateTodo)
		todo.PATCH("/:id", middlewares.AuthorizeToken(db, cfg), controller.UpdateTodo)
		todo.DELETE("/:id", middlewares.AuthorizeToken(db, cfg), controller.DeleteTodo)
	}
}
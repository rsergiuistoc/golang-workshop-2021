package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/models"
	"gorm.io/gorm"
	"net/http"
)

type TodoController struct {
	db 		*gorm.DB
}

func NewTodoController(d *gorm.DB) *TodoController{
	return &TodoController{
		db: d,
	}
}

func (t *TodoController) RetrieveTodo(c *gin.Context) {

	id := c.Param("id")
	var todo models.Todo

	err := t.db.Where("id = ?", id).First(&todo).Error
	if errors.Is(err, gorm.ErrRecordNotFound){
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (t *TodoController) CreateTodo(c *gin.Context) {

	user := c.MustGet("user").(models.User)

	var todo models.Todo

	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo.UserID = user.ID

	if err := t.db.Create(&todo).Error; err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func (t *TodoController) ListTodos(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	c.JSON(http.StatusOK, user.Todos)
}

func (t *TodoController) UpdateTodo(c *gin.Context) {
	var todo models.Todo

	id := c.Param("id")

	err := t.db.Where("id = ?", id).First(&todo).Error
	if errors.Is(err, gorm.ErrRecordNotFound){
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	err = c.ShouldBindJSON(&todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := t.db.Save(&todo).Error; err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, todo)
}

func (t *TodoController) DeleteTodo(c *gin.Context) {

	var todo models.Todo

	id := c.Param("id")

	err := t.db.Where("id = ?", id).Delete(&todo).Error
	if errors.Is(err, gorm.ErrRecordNotFound){
		c.JSON(http.StatusNotFound, "Todo Not Found")
		return
	}

	c.JSON(http.StatusNoContent, "Ok")
}

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

// RetrieveTodo godoc
// @Summary Retrieves a Todo item
// @Tags Todo
// @Produce json
// @Param id path string true "Id"
// @Success 200 {string} models.Todo
// @Failure 404 {string} Not Found
// @Router /todos/:id [get]
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

// CreateTodo godoc
// @Summary Creates a new Todo item
// @Tags Todo
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo Update" data
// @Success 200 {string} models.Todo
// @Failure 401 {string} Missing required fields
// @Router /todos/ [post]
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

// ListTodos godoc
// @Summary List all todos for the authenticated user.
// @Tags Todo
// @Success 200 {string} []models.Todo
// @Router /todos/ [get]
func (t *TodoController) ListTodos(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	c.JSON(http.StatusOK, user.Todos)
}

// UpdateTodo godoc
// @Summary Updates a Todo
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Param todo body models.Todo true "Todo Update" data
// @Success 200 {string} models.Todo
// @Failure 404 {string} Not Found
// @Router /todos/:id [patch]
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

// DeleteTodo godoc
// @Summary Updates a Todo
// @Tags Todo
// @Param id path string true "Id"
// @Success 204 {string} Ok
// @Failure 404 {string} Not Found
// @Router /todos/:id [delete]
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

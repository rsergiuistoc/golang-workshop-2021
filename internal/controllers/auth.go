package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/jwt"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/models"
	"gorm.io/gorm"
	"net/http"
)

type SingupRequestCommand struct {

	FirstName	string 	`json:"first_name"`
	LastName	string  `json:"last_name"`
	Email		string 	`json:"email"`
}

type AuthController struct{
	db 		*gorm.DB
	Config	*internal.Configuration
}

func NewAuthController(d *gorm.DB, cfg *internal.Configuration) *AuthController{
	return &AuthController{
		db: d,
		Config: cfg,
	}
}

func (a *AuthController)SignIn(c *gin.Context){

	user := c.MustGet("user").(*models.User)

	token, _ := jwt.EncodeToken(user.ID,  a.Config.SecretKey)

	c.JSON(http.StatusCreated, token)
	return
}

func (a *AuthController)SingUp(c *gin.Context){

	var data SingupRequestCommand

	err := c.ShouldBindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: data.FirstName,
		LastName: data.LastName,
		Email: data.Email,
	}

	if result := a.db.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusCreated, "Ok")
	return

}
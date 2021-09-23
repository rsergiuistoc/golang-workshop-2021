package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/jwt"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/models"
	"gorm.io/gorm"
	"net/http"
)

type SignupRequestCommand struct {

	FirstName	string 	`json:"first_name"`
	LastName	string  `json:"last_name"`
	Email		string 	`json:"email"`
	Password 	string 	`json:"password"`
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

// SignIn godoc
// @Summary Authenticates the User
// @Tags Authentication
// @Success 200 {string} Json Web Token
// @Failure 401 {string} Invalid email or password
// @Router /auth/signin [post]
// @securityDefinitions.basic BasicAuth
func (a *AuthController)SignIn(c *gin.Context){

	user := c.MustGet("user").(*models.User)

	token, _ := jwt.EncodeToken(user.ID,  a.Config.SecretKey)

	c.JSON(http.StatusCreated, token)
	return
}

// SignUp  godoc
// @Summary Registers the User
// @Tags Authentication
// @Accept json
// @Param signup body SignupRequestCommand true "Signup data"
// @Success 200 {string} Ok
// @Failure 401 {string} Missing required fields
// @Router /auth/signup [post]
func (a *AuthController) SignUp(c *gin.Context){

	var data SignupRequestCommand

	err := c.ShouldBindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: data.FirstName,
		LastName: data.LastName,
		Email: data.Email,
		Password: data.Password,
	}

	if result := a.db.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusCreated, "Ok")
	return

}
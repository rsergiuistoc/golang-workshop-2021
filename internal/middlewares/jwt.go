package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/jwt"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func AuthorizeToken(db *gorm.DB, cfg *internal.Configuration) gin.HandlerFunc {

	return func (c *gin.Context){
		auth := strings.Split(c.GetHeader("Authorization"), " ")

		if auth[0] != "Bearer"{
			failedAuthentication("Invalid authorization header.", c)
			return
		}

		if len(auth) == 1 {
			failedAuthentication("Invalid bearer header. No credentials provided.", c)
			return
		}

		if len(auth) > 2 {
			failedAuthentication("Invalid bearer header. Credentials string should not contain spaces.", c)
			return
		}

		claims, err := jwt.ValidateToken(auth[1], cfg.SecretKey)

		if err != nil{
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userId, _ := uuid.FromString(claims["user_id"].(string))

		var user models.User

		err = db.Preload("Todos").Where("id = ?", userId).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound){
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "User does not exist",
			})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
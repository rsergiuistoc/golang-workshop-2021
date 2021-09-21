package middlewares

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/models"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

// BasicAuthentication middleware against user credentials
func BasicAuthentication(db *gorm.DB) gin.HandlerFunc{

	return func (c *gin.Context){
		auth := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if auth[0] != "Basic"{
			failedAuthentication("Invalid authorization header.", c)
			return
		}

		if len(auth) == 1{
			failedAuthentication("Invalid basic header. No credentials provided.", c)
			return
		}

		if len(auth) > 2 {
			failedAuthentication("Invalid basic header. Credentials string should not contain spaces.", c)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(auth[1])

		if err != nil {
			failedAuthentication("Invalid basic header. Credentials not correctly base64 encoded.", c)
			return
		}

		decodedPayload := strings.Split(string(payload), ":")

		user := authenticateCredentials(decodedPayload[0], decodedPayload[1], db)
		if len(decodedPayload) != 2 || user == nil{
			failedAuthentication("Invalid email or password.", c)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func authenticateCredentials(email, password string, db *gorm.DB) *models.User {

	var user *models.User

	err := db.Where("email = ?", fmt.Sprintf("%s", email)).First(&user).Error
	if errors.Is(err , gorm.ErrRecordNotFound){
		return nil
	}

	err = models.CheckPassword(user.Password, password)
	if err != nil{
		return nil
	}
	return user
}

func failedAuthentication(message string, c *gin.Context){

	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
	c.Abort()
}
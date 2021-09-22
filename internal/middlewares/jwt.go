package middlewares

import (
	JWT "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rsergiuistoc/golang-workshop-2021/internal"
	"github.com/rsergiuistoc/golang-workshop-2021/internal/jwt"
	"net/http"
	"strings"
)

func AuthorizeToken(cfg *internal.Configuration) gin.HandlerFunc {

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

		token, err := jwt.ValidateToken(auth[1], cfg.SecretKey)

		if err != nil{
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token != nil && !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(JWT.MapClaims)

		c.Set("user", claims["user_id"])
		c.Next()
	}
}
package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthorizeToken() gin.HandlerFunc {

	return func (c *gin.Context){
		_ = strings.Split(c.Request.Header.Get("Authorization"), " ")
	}
}
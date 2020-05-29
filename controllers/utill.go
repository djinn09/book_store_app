package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware middleware gin server
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c)
		c.Next()
	}
}

// validateToken token validation
func validateToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	splitToken, err := StrngSplit(token, "")
	fmt.Println(splitToken[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized User"})
	}
	if len(splitToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized User"})
	}
	token = splitToken[1]
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized User"})
	} else if VerifyToken(token) {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized User"})
	}
}

// StrngSplit Split With separator
func StrngSplit(s string, separator string) ([]string, error) {
	var split []string
	if separator == "" {
		split = strings.Fields(s)
	} else {
		split = strings.Split(s, separator)
	}

	if len(split) == 0 {
		return []string{}, errors.New("String Not split")
	}
	return split, nil
}

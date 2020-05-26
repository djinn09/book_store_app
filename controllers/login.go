package controllers

import (
	"crud/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(userid uint) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(tokenString string) bool {
	var isValid = false
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return isValid
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"]
		var user models.Users
		if err := models.DB.Where("id=?", userID).First(&user).Error; err != nil {
			return isValid
		}
		return true
	}

	return isValid
}

/*
{
  "username": "Start with",
  "password": "Simon Sink"
}
*/
func LoginUser(c *gin.Context) {
	var u UserInput
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	var user models.Users
	c.Bind(&user)
	if err := models.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid Credentials"})
		return
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

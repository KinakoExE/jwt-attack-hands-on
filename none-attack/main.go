package main

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", top)
	r.GET("/token", token)
	r.GET("/admin", admin)

	r.Run(":5555")
}

var top = func(c *gin.Context) {
	c.JSON(200, gin.H{"message": "You need a token to access /admin. Please access /token"})
}

var token = func(c *gin.Context) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"user": "guest",
		"iat":  time.Now(),
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err == nil {
		c.JSON(200, gin.H{"token": tokenString})
	} else {
		c.JSON(500, gin.H{"message": "Something went wrong"})
	}
}

var admin = func(c *gin.Context) {

	if c.Request.Header["Authorization"] == nil {
		c.JSON(401, gin.H{"message": "Please set your JWT in Authorization header"})
		return
	}

	token, _ := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		b := jwt.UnsafeAllowNoneSignatureType // alg: noneを意図的に許可
		return b, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	if claims["user"] == "admin" {
		msg := fmt.Sprintf("Hello, %s !! Congrats this is JWT none attack!!", claims["user"])
		c.JSON(200, gin.H{"message": msg})
	} else {
		msg := fmt.Sprintf("Hello, %s , but you are not admin!", claims["user"])
		c.JSON(403, gin.H{"message": msg})
	}
}

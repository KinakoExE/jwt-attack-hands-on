package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/joho/godotenv"
)

type post struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()
	r.GET("/", top)
	r.GET("/token", token)
	r.GET("/admin", admin)

	r.Run(":5555")
}

var top = func(c *gin.Context) {
	c.JSON(200, gin.H{"message": "You need a token to access /admin. Please access /token."})
}

var token = func(c *gin.Context) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(".env file missing")
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = "guest"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
	if err == nil {
		c.JSON(200, gin.H{"token": tokenString})
	} else {
		c.JSON(500, gin.H{"message": "Something is wrong"})
	}
}

var admin = func(c *gin.Context) {

	if c.Request.Header["Authorization"] == nil {
		c.JSON(401, gin.H{"message": "Please set your JWT in Authorization header"})
		return
	}

	token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("SECRET"))
		return b, nil
	})

	if err != nil {
		c.JSON(500, gin.H{"message": "Something went wrong, please check whether your JWT format is valid or not"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["user"] == "admin" {
		msg := fmt.Sprintf("Hello %s !! Congrats you successfully exploited JWT by brute force!!", claims["user"])
		c.JSON(200, gin.H{"message": msg})
	} else {
		msg := fmt.Sprintf("Your username is %s ! You are not admin! Get out!!", claims["user"])
		c.JSON(200, gin.H{"message": msg})
	}
}

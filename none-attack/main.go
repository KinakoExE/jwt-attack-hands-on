package main

import (
	"fmt"
	"time"

	jwt "./jwt-go"
	"./jwt-go/request"
	"github.com/gin-gonic/gin"
)

const secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You need a token to access /admin. Plz acess /token"})

	})

	r.GET("/token", func(c *gin.Context) {
		token := jwt.New(jwt.GetSigningMethod("HS256"))

		token.Claims = jwt.MapClaims{
			"user": "guest",
			"exp":  time.Now().Add(time.Hour * 1).Unix(),
		}

		tokenString, err := token.SignedString([]byte(secretKey))
		if err == nil {
			c.JSON(200, gin.H{"token": tokenString})
		} else {
			c.JSON(500, gin.H{"message": "Could not generate token"})
		}
	})

	r.GET("/admin", func(c *gin.Context) {

		token, _ := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := jwt.UnsafeAllowNoneSignatureType
			return b, nil
		})
		claims := token.Claims.(jwt.MapClaims)
		if claims["user"] == "admin" {
			msg := fmt.Sprintf("Hello, %s !! Congratz You did JWT none attack", claims["user"])
			c.JSON(200, gin.H{"message": msg})
		} else {
			msg := fmt.Sprintf("Hello, %s , but you are not admin!", claims["user"])
			c.JSON(200, gin.H{"message": msg})
		}
	})

	r.Run(":5555")
}

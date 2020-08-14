package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"./auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
)

type post struct {
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()
	r.Handle("/", public)
	r.Handle("/token", auth.GetTokenHandler)
	r.Handle("/admin", auth.JwtMiddleware.Handler(admin))

	if err := http.ListenAndServe(":5555", r); err != nil {
		log.Fatalln("ListenAndServe:", nil)
	}
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Message: "You need a token to access /admin. Plz acess /token.",
	}
	json.NewEncoder(w).Encode(post)
})

var admin = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	token, _ := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("SIGNINGKEY"))
		return b, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	if claims["user"] == "admin" {
		post := &post{
			Message: "Congratz! You bruteforced JWT secret!",
		}
		json.NewEncoder(w).Encode(post)
	} else {
		msg := fmt.Sprintf("Your username is %s ! You are not admin! Get out!!", claims["username"])
		post := &post{
			Message: msg,
		}
		json.NewEncoder(w).Encode(post)
	}
})

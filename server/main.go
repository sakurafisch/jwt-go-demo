package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Super Secret Information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Header["Token"] != nil {
			token, err := jwt.Parse(req.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintln(res, err.Error())
			}
			if token.Valid {
				endpoint(res, req)
			}
		} else {
			fmt.Fprintf(res, "Not Authorized")
		}
	})
}

func handleRequest() {
	http.HandleFunc("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("My Simple Server")
	handleRequest()
}

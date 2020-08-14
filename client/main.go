package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// var mySigningKey = os.Get("MY_JWT_TOKEN")

var mySigningKey = []byte("mysupersecretphrase")

func homePage(res http.ResponseWriter, req *http.Request) {
	validToken, err := generateJWT()
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "HTTP://localhost:9000/", nil)
	request.Header.Set("Token", validToken)
	response, err := client.Do(request)
	if err != nil {
		fmt.Fprintf(res, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	fmt.Fprintf(res, string(body))
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "winnerwinter"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	handleRequests()
}

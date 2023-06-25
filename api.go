package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	AccessToken string `json:"access_token"`
}

func getToken() (error, Token) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		password     = os.Getenv("PASSWORD")
		username     = os.Getenv("USERNAME")
		clientSecret = os.Getenv("CLIENT_SECRET")
	)

	url := "https://api.whrcloud.com/oauth/token"
	method := "POST"

	payloadString := "password=" + password + "&client_secret=" + clientSecret + "&client_id=Whirlpool_Android&username=" + username + "&grant_type=password"

	payload := strings.NewReader(payloadString)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err, Token{}
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, Token{}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var response Response

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return err, Token{}
	}

	token := DecodeToken(response.AccessToken)

	token.AccessToken = response.AccessToken

	return nil, token
}

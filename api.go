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
	"sync"
)

var lock = &sync.Mutex{}

var tokenSingleton *Token

type Response struct {
	AccessToken string `json:"access_token"`
}

func getToken() (error, *Token) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		password     = os.Getenv("PASSWORD")
		username     = os.Getenv("USERNAME")
		clientSecret = os.Getenv("CLIENT_SECRET")
	)

	// If we already have a tokenSingleton, and it's not expired, return it
	if tokenSingleton != nil && !tokenSingleton.isExpired() {
		return nil, tokenSingleton
	}

	url := "https://api.whrcloud.com/oauth/token"
	method := "POST"

	payloadString := "password=" + password + "&client_secret=" + clientSecret + "&client_id=Whirlpool_Android&username=" + username + "&grant_type=password"

	payload := strings.NewReader(payloadString)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err, nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, nil
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
		return err, nil
	}

	newToken := DecodeToken(response.AccessToken)

	if tokenSingleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if tokenSingleton == nil {
			tokenSingleton = &newToken
			tokenSingleton.AccessToken = response.AccessToken
		}
	} else {
		lock.Lock()
		defer lock.Unlock()
		tokenSingleton.AccessToken = response.AccessToken
		tokenSingleton.Exp = newToken.Exp
	}
	return nil, tokenSingleton
}

func getData(appliance string) (error, ApplianceData) {
	err, token := getToken()
	if err != nil {
		fmt.Println(err)
		return err, ApplianceData{}
	}

	url := "https://api.whrcloud.com/api/v1/appliance/" + appliance
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err, ApplianceData{}
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err, ApplianceData{}
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
		return err, ApplianceData{}
	}

	// Unmarshal the JSON response into a struct
	var applianceData ApplianceData
	err = json.Unmarshal(body, &applianceData)
	if err != nil {
		fmt.Println(err)
		return err, ApplianceData{}
	}

	return nil, applianceData
}

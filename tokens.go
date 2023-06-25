package main

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"
)

type Token struct {
	AccountId   int      `json:"accountId"`
	RoleName    string   `json:"roleName"`
	CompanyId   int      `json:"companyId"`
	Username    string   `json:"UserName"`
	User_name   string   `json:"user_name"`
	Scope       []string `json:"scope"`
	Exp         int64    `json:"exp"`
	Authorities []string `json:"authorities"`
	Jti         string   `json:"jti"`
	ClientId    string   `json:"client_id"`
	SAID        []string `json:"SAID"`
	AccessToken string
}

func DecodeToken(token string) Token {
	payload := strings.Split(token, ".")[1]

	decoded, err := base64.RawURLEncoding.DecodeString(payload)

	if err != nil {
		panic(err)
	}

	var t Token
	err = json.Unmarshal(decoded, &t)
	if err != nil {
		panic(err)
	}

	return t
}

func (t *Token) isExpired() bool {
	return t.Exp < time.Now().Unix()
}

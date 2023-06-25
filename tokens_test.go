package main

import (
	"testing"
	"time"
)

func TestDecodeToken(t *testing.T) {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOjExMTExLCJSb2xlTmFtZSI6IlNUQU5EQVJEX1VTRVIiLCJjb21wYW55SWQiOjAsIlVzZXJOYW1lIjoiSm9obkRvZSIsInVzZXJfbmFtZSI6ImV4YW1wbGVAZ21haWwuY29tIiwic2NvcGUiOlsidHJ1c3QiLCJyZWFkIiwid3JpdGUiXSwiZXhwIjoxNjg3Njk1NTQ3LCJhdXRob3JpdGllcyI6WyJTVEFOREFSRF9VU0VSIl0sImp0aSI6IlhYWFhYWFgiLCJjbGllbnRfaWQiOiJXaGlybHBvb2xfQW5kcm9pZCIsIlNBSUQiOlsiWFhYWFhYWFhYIiwiWFhYWFhYWFhYIl19.eRVbTUqpppAOsGaoO0-WwE4dXNaiz0VUYkvCihU1W0c"

	decodedToken := DecodeToken(token)

	testToken := Token{
		AccountId:   11111,
		RoleName:    "STANDARD_USER",
		CompanyId:   0,
		Username:    "JohnDoe",
		User_name:   "example@gmail.com",
		Scope:       []string{"trust", "read", "write"},
		Exp:         1687695547,
		Authorities: []string{"STANDARD_USER"},
		Jti:         "XXXXXXX",
		ClientId:    "Whirlpool_Android",
		SAID: []string{
			"XXXXXXXXX",
			"XXXXXXXXX",
		},
	}

	if decodedToken.AccountId != testToken.AccountId {
		t.Errorf("Expected AccountId to be %d, got %d", testToken.AccountId, decodedToken.AccountId)
	}

	if decodedToken.RoleName != testToken.RoleName {
		t.Errorf("Expected RoleName to be %s, got %s", testToken.RoleName, decodedToken.RoleName)
	}

	if decodedToken.CompanyId != testToken.CompanyId {
		t.Errorf("Expected CompanyId to be %d, got %d", testToken.CompanyId, decodedToken.CompanyId)
	}

	if decodedToken.Username != testToken.Username {
		t.Errorf("Expected Username to be %s, got %s", testToken.Username, decodedToken.Username)
	}

	if decodedToken.User_name != testToken.User_name {
		t.Errorf("Expected User_name to be %s, got %s", testToken.User_name, decodedToken.User_name)
	}

	if decodedToken.Scope[0] != testToken.Scope[0] {
		t.Errorf("Expected Scope[0] to be %s, got %s", testToken.Scope[0], decodedToken.Scope[0])
	}

	if decodedToken.Scope[1] != testToken.Scope[1] {
		t.Errorf("Expected Scope[1] to be %s, got %s", testToken.Scope[1], decodedToken.Scope[1])
	}

	if decodedToken.Scope[2] != testToken.Scope[2] {
		t.Errorf("Expected Scope[2] to be %s, got %s", testToken.Scope[2], decodedToken.Scope[2])
	}

	if decodedToken.Exp != testToken.Exp {
		t.Errorf("Expected Exp to be %d, got %d", testToken.Exp, decodedToken.Exp)
	}

	if decodedToken.Authorities[0] != testToken.Authorities[0] {
		t.Errorf("Expected Authorities[0] to be %s, got %s", testToken.Authorities[0], decodedToken.Authorities[0])
	}

	if decodedToken.Jti != testToken.Jti {
		t.Errorf("Expected Jti to be %s, got %s", testToken.Jti, decodedToken.Jti)
	}

	if decodedToken.ClientId != testToken.ClientId {
		t.Errorf("Expected ClientId to be %s, got %s", testToken.ClientId, decodedToken.ClientId)
	}

	if decodedToken.SAID[0] != testToken.SAID[0] {
		t.Errorf("Expected SAID[0] to be %s, got %s", testToken.SAID[0], decodedToken.SAID[0])
	}

	if decodedToken.SAID[1] != testToken.SAID[1] {
		t.Errorf("Expected SAID[1] to be %s, got %s", testToken.SAID[1], decodedToken.SAID[1])
	}
}

func TestIsExpired(t *testing.T) {
	expiredToken := Token{
		Exp: time.Now().Unix() - 100,
	}

	if !expiredToken.isExpired() {
		t.Errorf("Expected tokenSingleton to be expired")
	}

	notExpiredToken := Token{
		Exp: time.Now().Unix() + 100,
	}

	if notExpiredToken.isExpired() {
		t.Errorf("Expected tokenSingleton to not be expired")
	}
}

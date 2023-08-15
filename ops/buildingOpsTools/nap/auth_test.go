package nap

import (
	"fmt"
	"testing"
)

func TestAuthToken(t *testing.T) {
	token := NewAuthToken("somerandomtoken")
	header := token.AuthorizationHeader()
	if header != "token somerandomtoken" {
		t.Fail()
	}
}

func TestAuthBasic(t *testing.T) {
	basic := NewAuthBasic("user", "passw0rd")
	header := basic.AuthorizationHeader()
	if header != "basic dXNlcjpwYXNzdzByZA==" {
		fmt.Println("header-> ", header, "actual-> ")
		t.Fail()
	}
}

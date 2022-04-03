package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("dtherkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	email := getCode("dtherhtun.cw@gmail.com")
	fmt.Println(email)
	email = getCode("dtherhtun@yahoo.com")
	fmt.Println(email)
}

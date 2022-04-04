package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "Et nostrud id nisi deserunt et. Deserunt nostrud qui commodo in nostrud aliquip est. Qui exercitation eiusmod sunt magna ullamco voluptate do pariatur sunt sunt eu ad consequat. Fugiat ex anim eiusmod est duis. Esse ad laboris esse consequat excepteur labore culpa ex. Nostrud amet minim fugiat elit reprehenderit adipisicing eiusmod minim cupidatat anim non. Minim enim occaecat magna enim pariatur pariatur consequat adipisicing in sint incididunt exercitation."
	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

package main

import (
	"encoding/base64"
	"fmt"
)

const encodeStd string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func main() {
	s := "Et nostrud id nisi deserunt et. Deserunt nostrud qui commodo in nostrud aliquip est. Qui exercitation eiusmod sunt magna ullamco voluptate do pariatur sunt sunt eu ad consequat. Fugiat ex anim eiusmod est duis. Esse ad laboris esse consequat excepteur labore culpa ex. Nostrud amet minim fugiat elit reprehenderit adipisicing eiusmod minim cupidatat anim non. Minim enim occaecat magna enim pariatur pariatur consequat adipisicing in sint incididunt exercitation."
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
}

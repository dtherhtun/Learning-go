package main

import (
	"encoding/json"
	"fmt"
)

type thumbnail struct {
	URL           string
	Height, Width int
}

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

func main() {
	var data img
	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

	for k, v := range data.IDs {
		fmt.Println(k, v)
	}

	fmt.Println(data.Thumbnail.URL)
}

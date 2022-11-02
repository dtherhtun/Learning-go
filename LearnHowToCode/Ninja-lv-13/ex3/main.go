package main

import (
	"fmt"

	"github.com/dtherhtun/Learning-go/Ninja-lv-13/ex3/quote"
	"github.com/dtherhtun/Learning-go/Ninja-lv-13/ex3/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	fmt.Println(word.UseCount("I love go go golang"))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}

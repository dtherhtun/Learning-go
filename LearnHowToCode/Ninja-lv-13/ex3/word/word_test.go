package word

import (
	"fmt"
	"testing"

	"github.com/dtherhtun/Learning-go/Ninja-lv-13/ex3/quote"
)

var result = map[string]int{"I": 1, "go": 2, "golang": 1, "love": 1}
var data string = "I love go go golang"

func TestUseCount(t *testing.T) {
	for d, v := range UseCount(data) {
		if v != result[d] {
			t.Error("Expected", result[d], "Get", v)
		}
	}
}

func TestCount(t *testing.T) {
	r := Count(data)
	c := 0
	for _, v := range result {
		c += v
	}
	if c != r {
		t.Error("Expected", c, "Get", r)
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount(data))
	// Output: map[I:1 go:2 golang:1 love:1]
}

func ExampleCount() {
	fmt.Println(Count(data))
	// Output: 5
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

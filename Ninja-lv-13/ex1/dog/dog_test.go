package dog

import (
	"fmt"
	"testing"
)

type testpairs struct {
	data   int
	result int
}

var test = []testpairs{
	{1, 7},
	{3, 21},
	{6, 42},
}

func TestYears(t *testing.T) {
	for _, v := range test {
		r := Years(v.data)
		if r != v.result {
			t.Error("Expected", v.result, "Get", r)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	for _, v := range test {
		r := YearsTwo(v.data)
		if r != v.result {
			t.Error("Expected", v.result, "Get", r)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(2))
	// Output: 14
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(3))
	// Output: 21
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(i)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(i)
	}
}

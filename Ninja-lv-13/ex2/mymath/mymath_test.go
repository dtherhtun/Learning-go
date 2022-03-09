package mymath

import (
	"fmt"
	"testing"
)

type testdata struct {
	data   []int
	result float64
}

var test = []testdata{
	{[]int{2, 4, 6}, 4},
	{[]int{1, 3, 5}, 3},
	{[]int{3, 6, 9, 12}, 7.5},
}

func TestCenteredAvg(t *testing.T) {
	for _, v := range test {
		r := CenteredAvg(v.data)
		if v.result != r {
			t.Error("Expected", v.result, "Get", r)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{3, 6, 9, 12}))
	// Output: 7.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range test {
			CenteredAvg(v.data)
		}
	}
}

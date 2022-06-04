// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

func main() {
	ar := []int32{50, 49, 38, 49, 78, 36, 25, 96, 10, 67, 78, 58, 98, 8, 53, 1, 4, 7, 29, 6, 59, 93, 74, 3, 67, 47, 12, 85, 84, 40, 81, 85, 89, 70, 33, 66, 6, 9, 13, 67, 75, 42, 24, 73, 49, 28, 25, 5, 86, 53, 10, 44, 45, 35, 47, 11, 81, 10, 47, 16, 49, 79, 52, 89, 100, 36, 6, 57, 96, 18, 23, 71, 11, 99, 95, 12, 78, 19, 16, 64, 23, 77, 7, 19, 11, 5, 81, 43, 14, 27, 11, 63, 57, 62, 3, 56, 50, 9, 13, 45}

	var count int32
	socks := make(map[int32]int32)
	for i := 0; i < len(ar); i++ {
		socks[ar[i]]++
	}
	fmt.Println(socks)
	for i, _ := range socks {
		fmt.Println("key->", i)
		fmt.Println("value->", socks[i])
		if socks[i]%2 == 0 {
			count += socks[i] / 2
			fmt.Println("%2=0->", socks[i]/2)
			fmt.Println("divided count->", count)
		} else {
			count += int32(math.Trunc(float64(socks[i] / 2)))
			fmt.Println("%2=point->", int32(math.Trunc(float64(socks[i]/2))))
			fmt.Println("float count->", count)
		}
	}
	fmt.Println(len(ar))

	fmt.Println(count)
}


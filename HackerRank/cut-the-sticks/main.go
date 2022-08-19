// hackerrank.com/challenges/cut-the-sticks/problem

package main

import "fmt"

func main() {
	arr := []int32{5, 4, 4, 2, 2, 8}
	resoult := cutTheSticks(arr)
	fmt.Println(resoult)

	arr2 := []int32{1, 2, 3, 4, 3, 3, 2, 1}
	fmt.Println(cutTheSticks(arr2))

}

func cutTheSticks(arr []int32) []int32 {
	// Write your code here
	var sticksCut []int32

	for true {
		min := arr[0]
		for _, v := range arr {
			if min > v {
				min = v
			}
		}
		var temp_arr []int32
		sticksCut = append(sticksCut, int32(len(arr)))
		for _, v := range arr {
			if v-min > 0 {
				temp_arr = append(temp_arr, v-min)
			}
		}
		if len(temp_arr) == 0 {
			break
		}
		arr = temp_arr
	}
	return sticksCut
}

package main

import "fmt"

func safeValue(vals []int, index int) (n int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	return vals[index], nil
}

func main() {
	vals := []int{1, 2, 3}
	// v := vals[10] will panic

	v, err := safeValue(vals, 10)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("v:", v)
}

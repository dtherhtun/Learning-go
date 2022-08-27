// https://www.hackerrank.com/challenges/halloween-sale/problem

package main

import "fmt"

func main() {
	p := int32(76)
	d := int32(90)
	m := int32(3)
	s := int32(4000)
	r := howManyGames(p, d, m, s)
	fmt.Println("result->", r)
}

func sum(data []int32) int32 {
	var sum int32
	for _, v := range data {
		sum += v
	}
	return sum
}

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	var data []int32
	a := p
	data = append(data, p)
	for a > m {
		if m > a-d {
			data = append(data, m)
			break
		}
		data = append(data, a-d)
		a -= d
	}
	for sum(data)+m <= s {
		data = append(data, m)
	}
	return int32(len(data))
}

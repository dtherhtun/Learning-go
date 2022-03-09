// Package word allow you to count word
package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count return the total world of paragraph
func Count(s string) int {
	data := strings.Split(s, " ")
	return len(data)
}

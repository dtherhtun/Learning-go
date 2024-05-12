package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readNumbers(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result [][]int
	var row []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}
		row = append(row, x)

		if len(row) == 20 {
			rowCopy := make([]int, len(row))
			copy(rowCopy, row)
			result = append(result, rowCopy)
			row = row[:0]
		}
	}

	return result
}

func product(x, y, dx, dy int, grid [][]int) int {
	product := 1
	for i := range 4 {
		nx, ny := x+(i*dx), y+(i*dy)
		if nx < 0 || nx >= 20 || ny < 0 || ny >= 20 {
			return 0
		}
		product *= grid[nx][ny]
	}
	return product

}

func main() {
	fd, _ := os.Open("./numbers.txt")
	grid := readNumbers(fd)
	directions := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}
	maxProduct := 0
	for i := range 20 {
		for j := range 20 {
			for _, direction := range directions {
				dx, dy := direction[0], direction[1]
				prod := product(i, j, dx, dy, grid)
				if prod > maxProduct {
					maxProduct = prod
					fmt.Println(i, j, dx, dy)
				}
			}
		}
	}
	fmt.Println(maxProduct)
	position(12, 6, 1, -1, grid)
}

func position(x, y, dx, dy int, grid [][]int) {
	for i := range 4 {
		nx, ny := x+(i*dx), y+(i*dy)
		fmt.Println(nx, ny, grid[nx][ny])
	}
}

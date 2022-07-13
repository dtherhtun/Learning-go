package caching

import "fmt"

const (
	rows = 4 * 1024
	cols = 4 * 1024
)

var metrix [rows][cols]byte

type data struct {
	v byte
	p *data
}

var list *data

func init() {
	var last *data

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var d data
			if list == nil {
				list = &d
			}
			if last != nil {
				last.p = &d
			}
			last = &d

			if row%2 == 0 {
				metrix[row][col] = 0xFF
				d.v = 0xFF
			}
		}
	}

	var ctr int
	d := list
	for d != nil {
		ctr++
		d = d.p
	}

	fmt.Println("Elements in the link list", ctr)
	fmt.Println("Elements in the matrix", rows*cols)
}

func LinkedListTraverse() int {
	var ctr int

	d := list
	for d != nil {
		if d.v == 0xFF {
			ctr++
		}
		d = d.p
	}

	return ctr
}

func ColumnTraverse() int {
	var ctr int

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if metrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

func RowTraverse() int {
	var ctr int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if metrix[row][col] == 0xFF {
				ctr++
			}
		}
	}

	return ctr
}

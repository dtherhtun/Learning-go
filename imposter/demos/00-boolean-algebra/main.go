package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func main() {
	data := [][]string{
		[]string{"1", "Bob", "52"},
		[]string{"2", "John", "34"},
		[]string{"3", "Jim", "83"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Name", "Age"})
	table.AppendBulk(data)

	table.Render()

	trueTable("AND", and)
	trueTable("OR", or)
	trueTable("NOT", not)
	trueTable("XOR", xor)
	trueTable("EQUIV", equiv)
	trueTable("IMP", imp)
	trueTable("NOR", nor)
	trueTable("NAND", nand)
	trueTable("ADDITION", add)
	trueTable("ADDLOGIC", addLogical)
}

type operation func(x, y int) string

func trueTable(op string, opFunc operation) {
	binaryData := [][]string{
		[]string{"0", "0", opFunc(0, 0)},
		[]string{"0", "1", opFunc(0, 1)},
		[]string{"1", "0", opFunc(1, 0)},
		[]string{"1", "1", opFunc(1, 1)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"X", "Y", op})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.AppendBulk(binaryData)
	table.Render()
}

func and(x, y int) string { return strconv.Itoa(x & y) }
func or(x, y int) string  { return strconv.Itoa(x | y) }
func not(x, y int) string { return strconv.Itoa(x &^ y) }
func xor(x, y int) string { return strconv.Itoa(x ^ y) }

func equiv(x, y int) string {
	if x == y {
		return strconv.Itoa(1)
	}
	return strconv.Itoa(0)
}

func imp(x, y int) string {
	if x == 1 {
		return strconv.Itoa(y)
	}
	return strconv.Itoa(1)
}

func nor(x, y int) string {
	if x != 1 && y != 1 {
		return strconv.Itoa(1)
	}
	return strconv.Itoa(0)
}

func nand(x, y int) string {
	if x != 1 || y != 1 {
		return strconv.Itoa(1)
	}
	return strconv.Itoa(0)
}

func add(x, y int) string {
	return fmt.Sprintf("%02b", x+y)
}

func addLogical(x, y int) string {
	return and(x, y) + xor(x, y)
}

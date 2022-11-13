package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
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
	trueTable("HalfAdder", halfAdder)
	trueTable("FullAddler", fullAdder)
	trueTable("HalfNadder", halfNadder)
	trueTable("FullAdder2", fullAdder2)
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

func halfAdder(x, y int) string {
	return and(x, y) + xor(x, y)
}
func halfNadder(x, y int) string {
	return or(x, y) + equiv(x, y)
}
func fullAdder(x, y int) string {
	carry := 1

	firstStep := halfAdder(x, y)
	firstRes0, _ := strconv.Atoi(firstStep[:1])
	firstRes1, _ := strconv.Atoi(firstStep[1:])
	secondStep := halfAdder(firstRes1, carry)
	secondRes0, _ := strconv.Atoi(secondStep[:1])
	secondRes1, _ := strconv.Atoi(secondStep[1:])

	leftResult := or(firstRes0, secondRes0)
	rightResult := secondRes1
	return leftResult + strconv.Itoa(rightResult)
}
func fullAdder2(x, y int) string {
	carry := 0
	if carry == 1 {
		return halfNadder(x, y)
	}
	return halfAdder(x, y)
}

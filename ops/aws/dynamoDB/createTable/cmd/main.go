package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/dtherhtun/Learning-go/ops/aws/dynamoDB/createTable"
)

const tableName = "books"

func main() {

	fmt.Printf("Creating table %v \n", tableName)

	if err := createTable.CreateTable(aws.String(tableName)); err != nil {
		fmt.Printf("Error creating table %v \n", err)
		os.Exit(1)
	}

	fmt.Printf("Waiting for table to exist")
	if err := createTable.Wait(aws.String(tableName)); err != nil {
		fmt.Printf("Error waiting for table to exist %v \n", err)
		os.Exit(1)
	}

	fmt.Printf("Table created \n")
	fmt.Printf("Fill table \n")
	if err := createTable.FillTable(aws.String(tableName)); err != nil {
		fmt.Printf("Error filling table %v \n", err)
		os.Exit(1)
	}
	fmt.Printf("Table filled \n")
}

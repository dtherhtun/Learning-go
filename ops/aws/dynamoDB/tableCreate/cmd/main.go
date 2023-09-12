package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/dtherhtun/Learning-go/ops/aws/dynamoDB/tableCreate"
)

const tableName = "barjokes"

func main() {

	fmt.Printf("Creating table %v \n", tableName)
	err := tableCreate.CreateTable(aws.String(tableName))
	if err != nil {
		fmt.Printf("Error creating table %v \n", err)
		os.Exit(1)
	}
	fmt.Printf("Waiting for table to exist")
	err = tableCreate.Wait(aws.String(tableName))
	if err != nil {
		fmt.Printf("Error waiting for table to exist %v \n", err)
		os.Exit(1)
	}
	fmt.Printf("Table created \n")
	fmt.Printf("Fill table \n")
	err = tableCreate.FillTable(aws.String(tableName))
	if err != nil {
		fmt.Printf("Error filling table %v \n", err)
		os.Exit(1)
	}
	fmt.Printf("Table filled \n")

}

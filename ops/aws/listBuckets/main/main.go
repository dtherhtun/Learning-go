package main

import (
	"fmt"

	"github.com/dtherhtun/Learning-go/ops/aws/listBuckets"
)

func main() {
	buckets, err := listBuckets.ListBuckets()
	if err != nil {
		fmt.Println(err)
	}

	for _, name := range buckets {
		fmt.Println(*name)
	}
}

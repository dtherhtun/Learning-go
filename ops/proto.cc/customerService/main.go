package main

import "github.com/dtherhtun/Learning-go/ops/proto.cc/customerService/customer"

func main() {
	a := customer.App{}
	a.Initialize()
	a.Run()
}

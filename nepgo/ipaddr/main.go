/*
IP
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: %s ip-addr\n", os.Args[0])
	}

	ip := os.Args[1]
	addr := net.ParseIP(ip)

	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", addr.String())
	}
}

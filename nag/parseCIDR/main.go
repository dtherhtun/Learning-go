package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	v4Addr, v4Net, err := net.ParseCIDR("192.168.0.1/24")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Addr-> %s, Net-> %s\n", v4Addr, v4Net)

	v6Addr, v6Net, err := net.ParseCIDR("2001:db8:a0b:12f0::1/32")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Addr-> %s, Net-> %s\n", v6Addr, v6Net)
}

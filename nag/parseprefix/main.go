package main

import (
	"fmt"
	"net/netip"
)

func main() {
	addr1 := "192.0.2.18"
	addr2 := "198.51.100.3"
	network4 := "192.0.2.0/24"

	pf := netip.MustParsePrefix(network4)
	fmt.Printf("prefix address: %v, length: %v\n", pf.Addr(), pf.Bits())

	ip1 := netip.MustParseAddr(addr1)
	if pf.Contains(ip1) {
		fmt.Println(addr1, "is in", network4)
	}

	ip2 := netip.MustParseAddr(addr2)
	if pf.Contains(ip2) {
		fmt.Println(addr2, "is in", network4)
	}
}

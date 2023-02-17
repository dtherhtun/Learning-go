package main

import (
	"fmt"
	"log"
	"net"
	"net/netip"
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

	IPv4, err := netip.ParseAddr("224.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	if IPv4.IsMulticast() {
		fmt.Println("IPv4 address is Multicast")
	}

	IPv6, err := netip.ParseAddr("FE80:F00D::1")
	if err != nil {
		log.Fatal(err)
	}

	if IPv6.IsLinkLocalUnicast() {
		fmt.Println("IPv6 address is link local Unicast")
	}

	v4 := net.ParseIP("192.0.2.1")
	if err != nil {
		log.Fatal(err)
	}
	v4s, _ := netip.AddrFromSlice(v4)
	fmt.Println(v4s.String())
	fmt.Println(v4s.Unmap().Is4())
}

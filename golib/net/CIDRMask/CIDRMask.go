package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// This mask corrsponds to a /31 subnet for IPv4.
	fmt.Println(net.CIDRMask(31, 32))

	// This mask corresponds to a 64 subnet for IPv6.
	fmt.Println(net.CIDRMask(64, 128))

	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.168.31.133/24")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("IPv4 Address:", ipv4Addr)
	fmt.Println("IPv4 Net:    ", ipv4Net)
}

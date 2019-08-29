package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("192.0.2.1")
	fmt.Println(ip.DefaultMask())

	ipv4Addr := net.ParseIP("192.168.31.133")
	fmt.Println(ipv4Addr)
	// This mask corresponds to a /24 subnet for IPv4
	ipv4Mask := net.CIDRMask(24, 32)
	fmt.Println(ipv4Addr.Mask(ipv4Mask))
}

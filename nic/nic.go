package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail to get interface addr due to: %v", err)
		os.Exit(-1)
	}

	for _, addr := range addrs {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "fail to parse CIDR due to: %v", err)
			os.Exit(-2)
		}

		fmt.Printf("local addr: %v\n", ip)
	}

}

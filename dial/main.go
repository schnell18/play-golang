package main

import (
	"fmt"
	"net"
)

func main() {
	dialer := net.Dialer{DualStack: true}
	conn, err := dialer.Dial("tcp", "ds.test-ipv6.com:80")
	if err != nil {
		panic("Dial")
	}
	fmt.Fprintf(conn, "GET /ip/ HTTP/1.0\r\n\r\n")
	buf := make([]byte, 4096)
	conn.Read(buf)
	println(string(buf))
}

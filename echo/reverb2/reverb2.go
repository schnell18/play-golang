package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// conn.SetDeadline(<-time.After(10 * time.Second))
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	c1 := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			c1 <- input.Text()
		}
	}()
	for {
		select {
		case <-time.After(10 * time.Second):
			return
		case msg := <-c1:
			go echo(c, msg, 1*time.Second)
		}
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

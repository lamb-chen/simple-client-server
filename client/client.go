package main

import (
	"fmt"
	"net"
)

func main() {
	msg := "Here is a message"
	conn, _ := net.Dial("tcp", "127.0.0.1:8083")
	fmt.Fprintln(conn, msg)
}

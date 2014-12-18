package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":6633")

	if err != nil {
		fmt.Print("error" + err.Error())
	} else {
		fmt.Fprintf(conn, "--PING--")
		conn.Close()
	}
}

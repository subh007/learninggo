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

		readbuff := make([]byte, 100)
		for {
			fmt.Fprintf(conn, "--PING--")
			byteCount, _ := conn.Read(readbuff)
			fmt.Print(string(readbuff[:byteCount]))
		}
	}
}

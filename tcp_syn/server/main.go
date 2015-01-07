package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9100")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer ln.Close()
	fmt.Println("server has started")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("received one connection")
		conn.close()
	}

}

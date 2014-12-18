package main

import (
	"fmt"
	"net"
)

// function to start the server
func startServer() {
	// start the tcp server and listen to port 6633.
	ln, err := net.Listen("tcp", ":6633")

	if err != nil {
		fmt.Print("error happened")
	} else {

		for {
			fmt.Print("server started")
			_, err := ln.Accept()

			if err != nil {
				fmt.Print("erro in channel")
			}

			fmt.Print("connection recived.")
		}
	}
}

func main() {
	startServer()
}

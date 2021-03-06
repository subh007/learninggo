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

		fmt.Print("start server in different thread")
		for {
			fmt.Print("server started")
			conn, err := ln.Accept()

			if err != nil {
				fmt.Print("erro in channel")
			}

			fmt.Print("connection recived.")

			// started the client handling in new thread.
			go func(conn net.Conn) {
				readbuff := make([]byte, 30)
				for {
					byteCount, err := conn.Read(readbuff)

					if err != nil {
						fmt.Print("socket is closed bye !!")
						conn.Close()
						break
					} else {
						fmt.Print(string(readbuff[:byteCount]))
						fmt.Fprint(conn, "--PONG--")
					}

				}
				fmt.Print("thread stop")
			}(conn)
		}
	}
}

func main() {
	startServer()
}

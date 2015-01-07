package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

func connect(address string, wg sync.WaitGroup) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer conn.Close()

	fmt.Println("connection is established")
	//wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		// wg.Add(1)
		fmt.Printf("connection %d", i)
		fmt.Println()
		go connect(":9100", wg)
	}

	wg.Wait()
	fmt.Print("all the connection served")
}

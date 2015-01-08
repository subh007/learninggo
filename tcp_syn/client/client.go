package main

import (
	"fmt"
	"net"
	// "os"
	"sync"
	"time"
)

func connect(address string, id int, wg *sync.WaitGroup) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err.Error())
		// os.Exit(-1)
	}
	defer conn.Close()

	fmt.Println("connection is established : ", id)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond)
		fmt.Println("connection", i)
		go connect(":9100", i, &wg)
	}

	wg.Wait()
}

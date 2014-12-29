// example
package main

import (
	"flag"
	"fmt"
	"github.com/subh007/goodl/go_ping_sweep"
)

func main() {
	host := flag.String("host", "", "ip address/ hostname to analyse the ping.")
	flag.Parse()

	fmt.Print(*host)

	if go_ping_sweep.IsAdmin() {
		fmt.Println("start the execution")
		go_ping_sweep.PingGoogle()
	}
}

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

		for i := 1; i < 10; i++ {
			res := go_ping_sweep.PingGoogle()
			fmt.Println("Ping time is taken -- " + res.TimePing)
		}
	}
}

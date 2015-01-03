// example
package main

import (
	"flag"
	"fmt"
	"github.com/subh007/goodl/go_ping_sweep"
	"os"
)

func main() {
	host := flag.String("host", "", "ip address/ hostname to analyse the ping.")
	flag.Parse()

	if host == nil {
		fmt.Println("usage: ./example --host <host name/ip>")
		os.Exit(-1)
	}

	if go_ping_sweep.IsAdmin() {
		for i := 1; i < 10; i++ {
			t := go_ping_sweep.PingAnalyse(*host)
			t.CreateTable()
		}
	}
}

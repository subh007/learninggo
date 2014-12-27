// example
package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "", "ip address/ hostname to analyse the ping.")
	flag.Parse()

	fmt.Print(*host)
}

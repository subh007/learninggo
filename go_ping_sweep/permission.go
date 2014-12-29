// permission
package go_ping_sweep

import (
	"fmt"
	"net"
	"os"
)

// This funciton return the admin priviledge for the
// executing.
func IsAdmin() bool {
	if os.Getuid() == 0 {
		return true
	} else {
		fmt.Println("must be run with the root priviledge.")
		os.Exit(-1)
	}
	return false
}

func PingGoogle() {
	_, err := net.Dial("ipv4:icmp", "127.0.0.1")
	if err != nil {
		fmt.Println("some bug")
	} else {
		fmt.Println("proceed with connection")
	}
}

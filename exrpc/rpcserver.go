// rcpserver.go project main.go
package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Args struct {
	Name     string
	Lastname string
}

// Only methods that satisfy these criteria will be made available for remote access; other methods will be ignored:
// - the method is exported.
// - the method has two arguments, both exported (or builtin) types.
// - the method's second argument is a pointer.
// - the method has return type error.
type Concater struct {
}

func (c *Concater) Concate(args *Args, reply *string) error {
	*reply = args.Name + args.Lastname
	//*reply = 5
	return nil
}

func main() {
	con := new(Concater)
	rpc.Register(con)

	listen, e := net.Listen("tcp", ":9090")
	if e != nil {
		fmt.Print("error!!")
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print("error !!")
		} else {
			fmt.Print("got the connection")
			go rpc.ServeConn(conn)
		}

	}
}

package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	Name     string
	Lastname string
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Print("error!")
	}

	args := Args{"sant", "singh"}
	var reply string

	err = client.Call("Concater.Concate", &args, &reply)

	if err != nil {
		fmt.Print("error!!")
		log.Fatal("error", err)
	} else {
		fmt.Print(reply)
	}
}

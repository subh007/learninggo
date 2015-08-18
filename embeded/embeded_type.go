package main

import (
	"fmt"
)

type Human struct {
	name string
}

type Robot struct {
	id string
	Human
}

func main() {
	fmt.Println("hello embedding")
	robot := Robot{
		id: "ANDROID",
	}

	robot.Human = Human{
		name: "Google",
	}

	fmt.Println("name :" + robot.Human.name)
}

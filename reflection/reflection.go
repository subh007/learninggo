// reflection
package main

import (
	"fmt"
	"reflect"
)

type Iterable interface {
}

func main() {
	list_1 := []int{1, 2, 3, 4}
	list_2 := []uint8{'a', 'b', 'c', 'd'}
	Map(list_1)
	Map(list_2)
}

func Map(list Iterable) {

	input := reflect.ValueOf(list)
	//_ := reflect.ValueOf(list)

	for i := 0; i < input.Len(); i++ {

		data_type := input.Index(i).Type().Name()
		if data_type == "int" {
			fmt.Print(input.Index(i).Int())
		} else if data_type == "uint8" {
			fmt.Printf("%c", input.Index(i).Uint())
		}
	}
}

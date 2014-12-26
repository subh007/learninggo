// reflection
package main

import (
	"fmt"
	//"reflect"
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

	//_ := reflect.ValueOf(list)

	//for i := 1; i < list.Len(); i++ {
	//	fmt.Print("go")
	//}

	for _, value := range list {
		fmt.Print(value)
	}
}

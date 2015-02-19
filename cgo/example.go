package main

/*
#include "test.h"
*/
import "C"
import "fmt"

func Example() {
	fmt.Println("this is go")
	fmt.Println(C.GoString(C.myprint(C.CString("go!!"))))
}

//export receiveC
func receiveC(msg *C.char) {
	fmt.Println(C.GoString(msg))
}

func main() {
	Example()
}

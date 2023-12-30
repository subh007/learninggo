package main

import (
	"fmt"
	"sync"
)

func search(wg *sync.WaitGroup, input []int, item int, start int, end int, sig chan<- bool) {
	if start <= end {
		if input[start] == item {
			fmt.Println("item found")
			sig <- true
		}
	}
	wg.Done()
}

func simple_search(input []int, item, start, end int) bool {
	if start <= end {
		if input[start] == item {
			fmt.Println("item found")
			return true
		}
	}
	return true
}
func main() {
	/*.sig := make(chan bool, 2)
	var wg sync.WaitGroup
	wg.Add(9)
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 9; i++ {
		go search(&wg, input, 10, i, len(input), sig)
	}

	wg.Wait()

	select {
	case <-sig:
		fmt.Println("number found")
	default:
		fmt.Println("not found")
	}*/

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 9; i++ {
		simple_search(input, 10, i, len(input))
	}
}

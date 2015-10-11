package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, err := os.Create("temp.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	w := bufio.NewWriter(outputFile)

	n, err := w.WriteString("test string")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("byte : %d", n)
	w.Flush()
	outputFile.Close()

}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func openFile(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	outputfile, err := os.Create("output.py")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer outputfile.Close()

	outBuffer := bufio.NewWriter(outputfile)
	var lineNumber int

	for {
		buffer, _, err = reader.ReadLine()
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
			return
		}
		if err == io.EOF {
			break
		}
		lineNumber++
		if lineNumber > 7 && lineNumber < 12 {
			outBuffer.WriteString("update line")
			continue
		}

		fmt.Println(string(buffer))
		nn, err := outBuffer.WriteString(string(buffer) + "\n")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("bytes wrote :%d line number :%d\n", nn, lineNumber)
		outBuffer.Flush()

	}
}

func main() {
	flag.Parse()
	openFile(flag.Arg(0))
}

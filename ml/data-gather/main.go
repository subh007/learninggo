package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("iris.data")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1
	rawCSVData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rawCSVData)
}

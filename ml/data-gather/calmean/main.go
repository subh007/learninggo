package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
)

func main() {
	irisFile, err := os.Open("iris.data")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	//create dataframe
	irisDF := dataframe.ReadCSV(irisFile)

	fmt.Println(irisDF)

	// retrieve column from dataframe
	sepalLength := irisDF.Col("sepal_length").Float()

	// calculate mean
	meanVal := stat.Mean(sepalLength, nil)

	// calculate mod
	mod, count := stat.Mode(sepalLength, nil)

	//calculate median
	median, err := stats.Median(sepalLength)

	fmt.Println(meanVal)
	fmt.Println(mod)
	fmt.Println(count)
	fmt.Println(median)

}

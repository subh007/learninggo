package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func main() {
	irisFile, err := os.Open("iris.csv")
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

	// calculate max
	max := floats.Max(sepalLength)

	// calculate min
	min := floats.Min(sepalLength)

	// calculate variance
	variance := stat.Variance(sepalLength, nil)

	// standered deviation
	stdDev := stat.StdDev(sepalLength, nil)

	// sort
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// calculate quantile
	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)

	fmt.Println(meanVal)
	fmt.Println(mod)
	fmt.Println(count)
	fmt.Println(median)

	fmt.Printf("max : %f", max)
	fmt.Printf("min : %f", min)
	fmt.Printf("variance : %f", variance)
	fmt.Printf("std deviation : %f\n", stdDev)
	fmt.Printf("quant25 : %f\n", quant25)
	fmt.Printf("quant25 : %f\n", quant50)
	fmt.Printf("quant25 : %f\n", quant75)

}

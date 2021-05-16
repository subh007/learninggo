package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/kniren/gota/dataframe"
	"github.com/sajari/regression"
)

func main() {
	file, err := os.Open("advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	advDF := dataframe.ReadCSV(file)

	trainNum := (4 * advDF.Nrow()) / 5
	testNum := advDF.Nrow() / 5

	if trainNum+testNum < advDF.Nrow() {
		trainNum++
	}

	trainIdx := make([]int, trainNum)
	testIdx := make([]int, testNum)

	fmt.Printf("train %d, test %d, total: %d", trainNum, testNum, advDF.Nrow())

	for i := 0; i < trainNum; i++ {
		trainIdx[i] = i
	}

	for i := 0; i < testNum; i++ {
		testIdx[i] = i + trainNum
	}

	trainingDF := advDF.Subset(trainIdx)
	testDF := advDF.Subset(testIdx)

	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	for idx, setName := range []string{"training.csv", "test.csv"} {
		file, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		w := bufio.NewWriter(file)
		setMap[idx].WriteCSV(w)
	}

	// training
	f, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4

	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// create regression data model
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")

	for idx, record := range trainingData {
		if idx == 0 {
			continue
		}

		yVal, _ := strconv.ParseFloat(record[3], 64)
		xVal, _ := strconv.ParseFloat(record[0], 64)

		// add regression data point
		r.Train(regression.DataPoint(yVal, []float64{xVal}))
	}

	// Train/fit the regression model
	r.Run()

	// Output the trained model Parameter
	fmt.Printf("\n regression formula: %v\n", r.Formula)

	// Prediction
	f1, err := os.Open("test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	testReader := csv.NewReader(f1)
	testReader.FieldsPerRecord = 4
	testdata, _ := testReader.ReadAll()

	var mAE float64
	for idx, record := range testdata {
		if idx == 0 {
			continue
		}

		yOb, _ := strconv.ParseFloat(record[3], 64)
		xOb, _ := strconv.ParseFloat(record[0], 64)

		yPredict, _ := r.Predict([]float64{xOb})

		fmt.Printf("TV adv: %f, obs Sales: %f, predicted sales: %f, error: %f\n", xOb, yOb, yPredict, math.Abs(yPredict-yOb))
		mAE += math.Abs(yPredict-yOb) / float64(len(testdata))
	}
	fmt.Printf("\nMAE: %0.2f\n", mAE)
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	file, err := os.Open("state_wise_wave_2.csv")
	if err != nil {
		log.Fatal("file read error")
	}
	defer file.Close()

	stateDF := dataframe.ReadCSV(file)

	//fmt.Println(distDF)

	cData := stateDF.Filter(
		dataframe.F{
			Colname:    "Status",
			Comparator: "==",
			Comparando: "Confirmed",
		},
	)

	confirmData := cData.Filter(
		dataframe.F{
			Colname:    "Date_YMD",
			Comparator: ">",
			Comparando: "2020-08-01",
		},
	)

	fmt.Println(confirmData)
	// // histogram for each col
	for _, colName := range confirmData.Names() {

		if colName == "BR" {
			fmt.Println(colName)

			v := make(plotter.Values, confirmData.Nrow())
			d := make([]float64, confirmData.Nrow())
			c, _ := confirmData.Col(colName).Int()

			for index, value := range c {
				v[index] = float64(value)
				d[index] = float64(value)
				//fmt.Printf("%d at %d\n", value, index)
			}

			// make a plot and set its title
			p := plot.New()
			p.Title.Text = "Histogram " + colName

			h, err := plotter.NewHist(v, 50)
			if err != nil {
				log.Fatal(err)
			}

			h.Normalize(1)

			p.Add(h)

			p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png")

			// calculate mean
			meanVal := stat.Mean(d, nil)

			// calculate mod
			mod, count := stat.Mode(d, nil)

			//calculate median
			median, err := stats.Median(d)

			// calculate max
			max := floats.Max(d)

			// calculate min
			min := floats.Min(d)

			// calculate variance
			variance := stat.Variance(d, nil)

			// standered deviation
			stdDev := stat.StdDev(d, nil)

			// sort
			inds := make([]int, len(v))
			floats.Argsort(v, inds)

			// calculate quantile
			quant25 := stat.Quantile(0.25, stat.Empirical, v, nil)
			quant50 := stat.Quantile(0.50, stat.Empirical, v, nil)
			quant75 := stat.Quantile(0.75, stat.Empirical, v, nil)

			fmt.Printf("mean: %f\n", meanVal)
			fmt.Printf("mod: %f\n", mod)
			fmt.Printf("count: %f\n", count)
			fmt.Printf("median: %f\n", median)

			fmt.Printf("max : %f\n", max)
			fmt.Printf("min : %f\n", min)
			fmt.Printf("variance : %f\n", variance)
			fmt.Printf("std deviation : %f\n", stdDev)
			fmt.Printf("quant25 : %f\n", quant25)
			fmt.Printf("quant50 : %f\n", quant50)
			fmt.Printf("quant75 : %f\n", quant75)
		}
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	file, err := os.Open("../iris.csv")
	if err != nil {
		log.Fatal("file read error")
	}
	defer file.Close()

	irisDF := dataframe.ReadCSV(file)

	// histogram for each col
	for _, colName := range irisDF.Names() {

		if colName != "species" {
			fmt.Println(colName)

			v := make(plotter.Values, irisDF.Nrow())

			for index, value := range irisDF.Col(colName).Float() {
				v[index] = value
			}

			// make a plot and set its title
			p := plot.New()
			p.Title.Text = "Histogram " + colName

			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			h.Normalize(1)

			p.Add(h)

			p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png")
		}
	}
}

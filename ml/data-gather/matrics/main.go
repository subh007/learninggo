package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	data := []float64{1.2, -2.0, 4, 5.0}
	a := mat.NewDense(2, 2, data)
	fa := mat.Formatted(a, mat.Prefix(" "))
	fmt.Printf("mat = %v\n\n", fa)
}

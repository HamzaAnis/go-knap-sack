package main

import (
	"fmt"
	"math"
)

var knap [][]float64

func knapSack(value []float64, weight []float64, cap int) float64 {
	if len(value) != len(weight) {
		fmt.Println("Invalid  size")
		return 0
	}
	return knapSackRecusive(value, weight, cap, len(value)-1)
}

func knapSackRecusive(value []float64, weight []float64, cap int, pos int) float64 {
	if pos < 0 {
		return 0
	}
	if weight[pos] > float64(cap) {
		return knapSackRecusive(value, weight, cap, pos-1)
	} else {
		A := math.Max(value[pos], knapSackRecusive(value, weight, cap-int(weight[pos]), pos-1))
		B := math.Max(value[pos], knapSackRecusive(value, weight, cap, pos-1))
		return math.Max(A, B)
	}
}
func main() {
	var v = []float64{60, 100, 120}
	var w = []float64{10, 20, 30}

	fmt.Println(knapSack(v, w, 50))
}

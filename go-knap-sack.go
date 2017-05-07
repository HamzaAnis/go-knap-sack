package main

import (
	"fmt"
	"sort"
)

type knap []knapsack

func (c knap) Len() int      { return len(c) }
func (c knap) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c knap) Less(i, j int) bool {
	return c[i].getValueP() > c[j].getValueP()
}

func knapSack(cap int, objects []knapsack) float64 {
	display(objects)
	sort.Sort(knap(objects))
	fmt.Print("After sorting\n")
	display(objects)
	return knapSackRecursive(cap, objects, 0)
}

func display(a []knapsack) {
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i].value, a[i].weight, a[i].getValueP(), "\n")
	}
}

func knapSackRecursive(cap int, objects []knapsack, pos int) float64 {
	if pos == len(objects) {
		return 0
	}
	if cap <= 0 {
		return float64(cap) * objects[pos-1].getValueP()
	}
	return objects[pos].value + knapSackRecursive(cap-int(objects[pos].weight), objects, pos+1)
}
func main() {
	objects := []knapsack{{60, 100}, {100, 20}, {120, 30}}
	fmt.Println(knapSack(50, objects))
}

type knapsack struct {
	value  float64
	weight float64
}

func (a knapsack) getValueP() float64 {
	return a.value / a.weight
}

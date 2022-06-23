package main

import (
	"math"
)

func Distance_Formula(x1 float64,y1 float64,x2 float64,y2 float64) int {
	x := math.Pow((x2 - x1), 2)
	y := math.Pow((y2 - y1), 2)
	d := math.Sqrt((x + y))
	return int(d);
}

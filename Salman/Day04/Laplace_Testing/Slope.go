package main

import (
  "math"
)
//--------- my function --mustafain
func slope_formula_2(a, b float64)float64{
  c := math.Pow(a+b, 2)
  return c

}

func slope(n1 float64, n2 float64) float64{
  return slope_formula_2(n1,n2)
  }
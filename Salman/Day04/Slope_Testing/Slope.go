package main

import (
  "math"
)
//--------- my function --mustafain
func Slope_formula_2(a float64, b float64)float64{
  c := math.Pow(a+b, 2)
  return c

}

func Slope(n1 float64, n2 float64) float64{
  return Slope_formula_2(n1,n2)
  }
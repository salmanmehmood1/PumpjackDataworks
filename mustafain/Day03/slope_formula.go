package main

import (
  "fmt"
  "math"
)

var choice int
//--------- my function --mustafain
func slope_formula_2(a, b float64)float64{
  c := math.Pow(a+b, 2)
  return c

}

func slope(){
  var n1 float64
  var n2 float64
  fmt.Println("Enter Your First number: ")
  fmt.Scanln(&n1)
  fmt.Println("Enter Second number: ")
  fmt.Scanln(&n2)
  fmt.Printf("Result : %.2f", slope_formula_2(n1,n2))
  }
//---------


func main() {
	slope()
}

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

fmt.Println("Enter your choice in number: ")
fmt.Scanln(&choice)


switch choice {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    slope()
case 4:
    fmt.Println("Thursday")
case 5:
    fmt.Println("Friday")
case 6:
    fmt.Println("Saturday")
case 7:
    fmt.Println("Sunday")
default:
    fmt.Println("Invalid")
}
}
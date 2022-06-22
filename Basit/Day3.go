package main

import "fmt"

func Slope_Formula(x11 float32,x22 float32,y11 float32,y22 float32) float32 {
    var gradient float32
    gradient=(y22-y11)/(x22-x11)
    return gradient
}
func slope1() {
  
    var x1,x2,y1,y2 float32
    var output float32
    fmt.Println("Enter the first xcoordinate: ")
    fmt.Scanln(&x1)
    fmt.Println("Enter Second xcoordinate: ")
    fmt.Scanln(&x2)
    fmt.Println("Enter the first ycoordinate: ")
    fmt.Scanln(&y1)
    fmt.Println("Enter Second ycoordinate: ")
    fmt.Scanln(&y2)
    output=Slope_Formula(x1,x2,y1,y2)
    
    fmt.Print("The slope of the given formula is:",output)
}
func main(){
    slope1()
}





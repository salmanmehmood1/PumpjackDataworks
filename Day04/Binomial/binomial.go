package test

import (
	"fmt"
	"math"
)

// func main() {
// 	Binomial()
// }

func factorial(n float64) float64 {
	var factVal float64 = 1 // float64 is the set of all unsigned 64-bit integers.

	if n < 0 {
		fmt.Print("Factorial of negative number doesn't exist.")
	} else {
		for i := 1; i <= int(n); i++ {
			factVal *= float64(i) // mismatched types int64 and int
		}

	}
	return factVal /* return from function*/
}

func Binomial(x float64, y float64) float64{

	var bino float64
	bino = math.Floor(factorial(x) / factorial(y) / factorial(x-y))

	fmt.Println("Binomial Coefficient: ", bino)
	return bino;

	// Logic:
	// flor( fac(x) / fac(y) / fac(x - y) )
}

func B() {
	fmt.Println()
	fmt.Println("Binomial Coefficient Calculator.")
	fmt.Println("================================")
	fmt.Println()

	var x float64
	var y float64
	fmt.Println("Enter First Integer: ")
	fmt.Scanln(&x)
	fmt.Println("Enter Second Integer: ")
	fmt.Scanln(&y)
	Binomial(x, y)

}

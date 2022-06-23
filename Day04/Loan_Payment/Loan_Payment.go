package Loan_Payment

import (
	"fmt"
	"math"
)
func loan(p float64,r float64,n float64,) float64 {
	var n1 float64
	var x float64
	var y float64
	var a float64
	 n1 = n - 1
	 x = (1 + r)
	 y = math.Pow(x, n)
	x = math.Pow(x, n1)
	a = p * ((r * (y)) / x)
	return a
}
func LoanPayment() {
	var p float64
	var r float64
	var n float64
	var output float64
	fmt.Println("Enter your loan amount : ")
	fmt.Scanln(&p)
	fmt.Println("Enter interest rate per period : ")
	fmt.Scanln(&r)
	fmt.Println("Enter total number of period: ")
	fmt.Scanln(&n)
	output=loan(p,r,n)
	fmt.Println("Payment amount per period : ", output)
}

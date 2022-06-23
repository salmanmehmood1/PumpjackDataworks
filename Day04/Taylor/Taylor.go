package Taylor

import (
	"fmt"
	"math"
)

func TaylorExpansion(x float64 ,num float64)float64 {
	var fact float64 = 1.0
	var sum float64 = 0.0
	var i float64

	for i = 1.0; i < num; i++ {
		fact = fact * i
		sum = sum + (math.Pow(x, i) / fact)
	}
	sum+=1;

	return sum;
}

func Taylor() {
	var x float64 = 2.0
	var num float64 = 10.0
	
	fmt.Println("Enter value of x: ");
	fmt.Scanln(&x);
	fmt.Println("Enter number of terms in taylor series: ");
	fmt.Scanln(&num);
	fmt.Println("The value of Tayloe Expansion is:",TaylorExpansion(x,num))
}

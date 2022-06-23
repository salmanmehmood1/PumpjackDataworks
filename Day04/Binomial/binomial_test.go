package Binomial

import (
	"fmt"
	"testing"
)

type binomialTest struct {
	arg1, arg2, expected float64
}

var binomialTests = []binomialTest{
	//1
	binomialTest{8, 3, 56},
	//2
	binomialTest{-9, 9, 0},
	//3
	binomialTest{999, 999, 1},
}

func TestBinomial(t *testing.T) {
	fmt.Println()
	fmt.Println("---------Test Case(TBinomial)--------")
	i := 1
	for _, test := range binomialTests {
		if output := Binomial(test.arg1, test.arg2); output != test.expected {
			fmt.Println("failed : Test case = ", i)
			t.Errorf("test failed, expected : '%f', got: '%f' ", output, test.expected)
		}
		i += 1
	}
	fmt.Println("--------------------------------------")
	fmt.Println()

}

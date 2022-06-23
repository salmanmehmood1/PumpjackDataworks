package Taylor

import (
	"fmt"
	"testing"
)

type taylorTest struct {
	arg1, arg2, expected float64
}

var taylorTests = []taylorTest{
	//1
	taylorTest{2, 10, 7.3887125220458545},
	//2
	taylorTest{0, 10, 1.0},
	//3
	taylorTest{-10, 10, -27.70631023742594},
}

func TestTaylor(t *testing.T) {
	fmt.Println();
	fmt.Println("---------Test Case(Taylor Series)--------")
	i := 1
	for _, test := range taylorTests {
		if output := TaylorExpansion(test.arg1, test.arg2); output != test.expected {
			fmt.Println("failed : Test case = ", i)
			t.Errorf("test failed, expected : '%.2f', got: '%.2f' ", output, test.expected)
		}
		i += 1
	}
	fmt.Println("--------------------------------------")
	fmt.Println();

}

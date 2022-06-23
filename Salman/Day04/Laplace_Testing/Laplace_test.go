package main

import (
    "testing"
)
type laplace_test struct{
	arg1,arg2,arg3 int
	expected float64
}
var addTests=[]laplace_test{
	laplace_test{23,23,50,3},
	laplace_test{59,59,80,30},
	laplace_test{10,10,100,6},
}
func TestLaplace(t *testing.T){
	for _, test:= range addTests{
		if output:=  Laplace(test.arg1,test.arg2,test.arg3); output!= test.expected{
			t.Errorf("Output %f not equal to expected %f",output,test.expected)
		}
	}
}
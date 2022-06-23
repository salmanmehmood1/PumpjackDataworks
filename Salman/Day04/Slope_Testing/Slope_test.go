package main

import (
    "testing"
)
type Slope_test struct{
	arg1,arg2 float64
	expected float64
}
var addTests=[]Slope_test{
	Slope_test{1,2,6},
	Slope_test{2,1,8},
	Slope_test{1,2,7},
}
func TestSlope(t *testing.T){
	for _, test:= range addTests{
		if output:=  Slope(test.arg1,test.arg2); output!= test.expected{
			t.Errorf("Output %f not equal to expected %f",output,test.expected)
		}
	}
}
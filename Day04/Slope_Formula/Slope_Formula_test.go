package main

import (
    "testing"
    "math"
)
type slopetest struct{
	arg1,arg2,arg3,arg4,expected float64
}
var addTests=[]slopetest{
	slopetest{1,2,3,5,2},
	slopetest{2,1,10,30,-20},
	slopetest{50,50,200,100,math.Inf(-1)},  // This test case should give error as denominator is zero so answer will be infinite where as expecrted output is -10
	slopetest{100,200,300,400,1},
}
func TestSlope1(t *testing.T){
	for _, test:= range addTests{
		if output:= Slope_Formula(test.arg1,test.arg2,test.arg3,test.arg4); output!= test.expected{
			t.Errorf("Output %f not equal to expected %f",output,test.expected)
		}
	}
}

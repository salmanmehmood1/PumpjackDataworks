package main

import (
    "testing"
)
type distancetest struct{
	arg1,arg2,arg3,arg4 float64
	expected int
}
var addTests=[]distancetest{
	distancetest{1,2,3,5,3},
	distancetest{2,1,10,30,30},
	distancetest{1,2,5,6,5},
}
func TestDistance(t *testing.T){
	for _, test:= range addTests{
		if output:= Distance_Formula(test.arg1,test.arg2,test.arg3,test.arg4); output!= test.expected{
			t.Errorf("Output %d not equal to expected %d",output,test.expected)
		}
	}
}

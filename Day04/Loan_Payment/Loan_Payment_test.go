package Loan_Payment

import (
	"fmt"
	"testing"
)
type LoanTest struct{
	arg1,arg2,arg3,expected float64
}
var LoanTests=[]LoanTest{
	LoanTest{100000, 2, 12,66666.67},
	LoanTest{10000,0,0,0.00},
	LoanTest{1234,3,0,925}, 
	LoanTest{13456,0,4,0.00},
}
func TestLoan(t *testing.T){
	i:=1

	for _, test:= range LoanTests{
		if output:= loan(test.arg1,test.arg2,test.arg3); output!= test.expected{
			fmt.Println("Test Case :",i," Failed")
			t.Errorf("Output %f not equal to expected %f",output,test.expected)
		}
		i+=1;
	}
}

package main

//function

//end function


import "testing"

func TestReverse(t *testing.T) {
    cases := []struct {
        input          int
        expectedOutput int
    }{
        {0, 0},
	{1, 1},
    }

    for _, c := range cases {
        if output := fourier_series(c.input); output != c.expectedOutput {
            t.Errorf("incorrect output for `%d`: expected `%d` but got `%d`", c.input, c.expectedOutput, output)
        }
    }
}
package main

import (
	"fmt"
  "math"
	"math/cmplx"
)

type Block struct {
    Try     func()
    Catch   func(Exception)
    Finally func()
}
 
type Exception interface{}
 
func Throw(up Exception) {
    panic(up)
}
func (tcf Block) Do() {
    if tcf.Finally != nil {
 
        defer tcf.Finally()
    }
    if tcf.Catch != nil {
        defer func() {
            if r := recover(); r != nil {
                tcf.Catch(r)
            }
        }()
    }
    tcf.Try()
}
//import libraries as not importing so copied
func fft(a []complex128, n int) []complex128 {
	x := make([]complex128, n)
	copy(x, a)

	j := 0
	for i := 0; i < n; i++ {
		if i < j {
			x[i], x[j] = x[j], x[i]
		}
		m := n / 2
		for {
			if j < m {
				break
			}
			j = j - m
			m = m / 2
			if m < 2 {
				break
			}
		}
		j = j + m
	}
	kmax := 1
	for {
		if kmax >= n {
			return x
		}
		istep := kmax * 2
		for k := 0; k < kmax; k++ {
			theta := complex(0.0, -1.0*math.Pi*float64(k)/float64(kmax))
			for i := k; i < n; i += istep {
				j := i + kmax
				temp := x[j] * cmplx.Exp(theta)
				x[j] = x[i] - temp
				x[i] = x[i] + temp
			}
		}
		kmax = istep
	}
}
func FFT(x []complex128, n int) []complex128 {
	y := fft(x, n)
	for i := 0; i < n; i++ {
		y[i] = y[i] / complex(float64(n), 0.0)
	}
	return y
}


func fourier_series(na int) int{
x0 := []float64{
		5,
		32,
		38,
		-33,
		-19,
		-10,
		1,
		-8,
		-20,
		10,
		-1,
		4,
		11,
		-1,
		-7,
		-2,
	}
  n := len(x0)
	x := make([]complex128, n)
	for k := 0; k < n; k++ {
		x[k] = complex(x0[k], 0.0)
	}

  y := FFT(x, n)
	Block{
        Try: func() {
            y := FFT(x, n)
		if(na==0){y=nil}
	if(y==nil){
            Throw("Oh,...sh...")
	}
        },
        Catch: func(e Exception) {
            fmt.Printf("Caught %v\n", e)
        },
        Finally: func() {
            fmt.Println("Finally...")
        },
    }.Do()
	
	

	fmt.Println(" K   DATA  FOURIER TRANSFORM  ")
	for k := 0; k < n ; k++ {
		fmt.Printf("%2d %6.1f  %8.3f%8.3f   \n",
			k, x0[k], real(y[k]), imag(y[k]))
	}
if(na==0){
return 0;
}
return 1;
  }

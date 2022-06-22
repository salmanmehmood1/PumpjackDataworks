package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Abdul Rehman



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


func fourier_series(){
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
  }
//end Abdul rehman
var choice int
//--------- my function --mustafain
func slope_formula_2(a, b float64)float64{
  c := math.Pow(a+b, 2)
  return c

}

func slope(){
  var n1 float64
  var n2 float64
  fmt.Println("Enter Your First number: ")
  fmt.Scanln(&n1)
  fmt.Println("Enter Second number: ")
  fmt.Scanln(&n2)
  fmt.Printf("Result : %.2f", slope_formula_2(n1,n2))
  }
//---------
//Basit Code Part

func Slope_Formula(x11 float32,x22 float32,y11 float32,y22 float32) float32 {
    var gradient float32
    gradient=(y22-y11)/(x22-x11)
    return gradient
}
func slope1() {
  
    var x1,x2,y1,y2 float32
    var output float32
    fmt.Println("Enter the first xcoordinate: ")
    fmt.Scanln(&x1)
    fmt.Println("Enter Second xcoordinate: ")
    fmt.Scanln(&x2)
    fmt.Println("Enter the first ycoordinate: ")
    fmt.Scanln(&y1)
    fmt.Println("Enter Second ycoordinate: ")
    fmt.Scanln(&y2)
    output=Slope_Formula(x1,x2,y1,y2)
    
    fmt.Print("The slope of the given formula is:",output)
}

 
// anabia

func distance() {
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64
	fmt.Print("Enter x coordinate of 1st point: ")
	fmt.Scan(&x1)
	fmt.Print("Enter y coordinate of 1st point: ")
	fmt.Scan(&y1)
	fmt.Print("Enter x coordinate of 2nd point: ")
	fmt.Scan(&x2)
	fmt.Print("Enter y coordinate of 2nd point: ")
	fmt.Scan(&y2)
	x := math.Pow((x2 - x1), 2)
	y := math.Pow((y2 - y1), 2)
	d := math.Sqrt((x + y))
	fmt.Print("distance is ", d)
}

//anabia

//erum
func quadraticformula() {
var a, b, c, root1, root2, imaginary, discriminant float64

    fmt.Print("Enter the a, b, c of Quadratic equation = ")
    fmt.Scanln(&a, &b, &c)

    discriminant = (b * b) - (4 * a * c)

    if discriminant > 0 {
        root1 = (-b + math.Sqrt(discriminant)/(2*a))
        root2 = (-b - math.Sqrt(discriminant)/(2*a))
        fmt.Println("Two Distinct Real Roots Exist: root1 = ", root1, " and root2 = ", root2)
    } else if discriminant == 0 {
        root1 = -b / (2 * a)
        root2 = -b / (2 * a)
        fmt.Println("Two Equal and Real Roots Exist: root1 = ", root1, " and root2 = ", root2)
    } else if discriminant < 0 {
        root1 = -b / (2 * a)
        root2 = -b / (2 * a)
        imaginary = math.Sqrt(-discriminant) / (2 * a)
        fmt.Println("Two Distinct Complex Roots Exist: root1 = ", root1, "+", imaginary, " and root2 = ", root2, "-", imaginary)
    }
}
//erum
func main() {

fmt.Println("Enter your choice in number: ")
fmt.Scanln(&choice)


switch choice {
case 1:
	fourier_series()
case 2:
	distance()
case 3:
    slope()
case 4:
    quadraticformula()
case 5:
    slope1()
default:
    fmt.Println("Invalid")
}
}

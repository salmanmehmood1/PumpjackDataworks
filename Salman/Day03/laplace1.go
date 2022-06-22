package main

import (
	"fmt"
)

func Laplace(Nx int, Ny int, iter int) {
	var temp, diff float32
	if Nx == Ny {
		for i := 1; i < Ny-1; i++ {
			for j := 1; j < Ny-1; j++ {
				temp = 0.25 * float32(i*j)
				diff = (temp - float32(Nx)) / 100
			}
		}
	}
	iter = iter - 1
	fmt.Print("Laplace Output : ", diff)

}

func main() {
	var Nx, Ny, iter int
	fmt.Print("Enter Nx: ")
	fmt.Scan(&Nx)
	fmt.Print("Enter Ny: ")
	fmt.Scan(&Ny)
	fmt.Print("Enter iteration No: ")
	fmt.Scan(&iter)
	Laplace(Nx, Ny, iter)
}

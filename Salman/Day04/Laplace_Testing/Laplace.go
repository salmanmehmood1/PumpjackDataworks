package main

func Laplace(Nx int, Ny int, iter int) float64{
	var temp, diff float64
	if Nx == Ny {
		for i := 1; i < Ny-1; i++ {
			for j := 1; j < Ny-1; j++ {
				temp = 0.25 * float64(i*j)
				diff = (temp - float64(Nx)) / 100
			}
		}
	}
	iter = iter - 1
	// fmt.Print("Laplace Output : ", diff)
  return diff;

}
######### Basit Code Part
def slope_formula(x11,x22,y11,y22):
    gradient= (y22-y11)/(x22-x11)
    return gradient
#########

import numpy, math
import sys
def relax(A, maxsteps, convergence):
    iterations = 0
    diff = convergence +1

    Nx = A.shape[1]
    Ny = A.shape[0]
    
    while iterations < maxsteps and diff > convergence:
        Atemp = A.copy()
        diff = 0.0
        
        for y in range(1,Ny-1):
            for x in range(1,Ny-1):
                A[y,x] = 0.25*(Atemp[y,x+1]+Atemp[y,x-1]+Atemp[y+1,x]+Atemp[y-1,x])
                diff  += math.fabs(A[y,x] - Atemp[y,x])

        diff /=(Nx*Ny)
        iterations += 1
    print ("Output : ", diff)

def boundary(A,x,y):


    Nx = A.shape[1]
    Ny = A.shape[0]
    Lx = x[Nx-1] 
    Ly = x[Nx-1]


    A[:,0]    =   100*numpy.sin(math.pi*x/Lx)
    A[:,Nx-1] = - 100*numpy.sin(math.pi*x/Lx)
    A[0,:]    = 0.0
    A[Ny-1,:] = 0.0

def fourierseries():
    return "monday"
def slopeformula1():
    x1=float(input("Enter the first xcoordinate:"))
    x2=float(input("Enter the second xcoordinate:"))
    y1=float(input("Enter the first ycoordinate:"))
    y2=float(input("Enter the second ycoordinate:"))
    Output=slope_formula(x1,x2,y1,y2)
    print("The slope of the given formula is:",Output)
def slopeformula2():
    return "wednesday"
def quadraticformula():
    return "thursday"
def laplaceformula():
    Nx = input("Enter Nx : ")
    Ny = input("Enter Ny : ")
    if Nx!=Ny:
        print("Nx and Ny must be equal")
        exit()
    maxiter = input("Enter no of iterations (P) : ")
    Nx=int(Nx)
    Ny=int(Ny)
    maxiter=int(maxiter)

    x = numpy.linspace(0,1,num=Nx+2)
    y = numpy.linspace(0,1,num=Ny+2)
    A = numpy.zeros((Nx+2,Ny+2))

    boundary(A,x,y)
    relax(A,maxiter,0.00001)
    # return "friday"
def default():
    return "Invalid choice"

switcher = {
    1: fourierseries,
    2: slopeformula1,
    3: slopeformula2,
    4: quadraticformula,
    5: laplaceformula
    }
  
def switch(formula):
    return switcher.get(formula, default)()

if __name__ == "__main__":
    formula = float(input('Enter choice: '))
    print(switch(formula))
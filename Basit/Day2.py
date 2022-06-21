x1=float(input("Enter the first xcoordinate:"))
x2=float(input("Enter the second xcoordinate:"))
y1=float(input("Enter the first ycoordinate:"))
y2=float(input("Enter the second ycoordinate:"))
def slope_formula(x11,x22,y11,y22):
    gradient= (y22-y11)/(x22-x11)
    return gradient
Output=slope_formula(x1,x2,y1,y2)
print("The slope of the given formula is:",Output)

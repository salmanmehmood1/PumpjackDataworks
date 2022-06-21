######### Basit Code Part
def slope_formula(x11,x22,y11,y22):
    gradient= (y22-y11)/(x22-x11)
    return gradient
#########

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
    return "friday"
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
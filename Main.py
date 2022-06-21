def fourierseries():
    return "monday"
def slopeformula1():
    return "tuesday"
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
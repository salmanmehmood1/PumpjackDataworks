import cmath  
a = float(input('Enter a: '))  
b = float(input('Enter b: '))  
c = float(input('Enter c: '))  
 
determinant = (b**2) - (4*a*c)  

if determinant > 0:
      root1 = (-b-cmath.sqrt(d))/(2*a) 
      root2 = (-b+cmath.sqrt(d))/(2*a) 

      print('root1 = {0} and root2 = {1}'.format(root1,root2))

elif determinant == 0:
      root1 = root2 = -b / (2 * a)
      print('root1 = root2 =', root1)
    

else:
      real = -b/(2*a)
      imaginary = cmath.sqrt(-determinant)/(2*a)
      print('root1 = {0}+{1}i'.format(real,imaginary))
      print('root2 = {0}+{1}i'.format(real,imaginary))

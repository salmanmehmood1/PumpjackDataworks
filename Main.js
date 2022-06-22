function quadEquationSolver(a, b, c) {
  var root1, root2;
  var determinant = b * b - 4 * a * c;

  if (determinant > 0) {
    root1 = (-b + Math.sqrt(determinant)) / (2 * a);
    root2 = (-b - Math.sqrt(determinant)) / (2 * a);

    console.log("root1 = %.2f and root2 = %.2f", root1, root2);
  }

  else if (determinant == 0) {
    root1 = root2 = -b / (2 * a);
    console.log("root1 = root2 = %.2f;", root1);
  }

  else {
    var real = -b / (2 * a);
    var imaginary = Math.sqrt(-determinant) / (2 * a);
    console.log("root1 = %.2f+%.2fi", real, imaginary);
    console.log("\nroot2 = %.2f-%.2fi", real, imaginary);
  }

};

function slope_formula(x11,x22,y11,y22){
  let gradient= (y2-y1)/(x2-x1);
  return gradient;
};

function square(d, e) {
  let c = Math.pow(d + e, 2);
  return c;
}

document.write("Choose the formula. \n");
document.write("1-fourierseries \n");
document.write("2-slopeformula1 \n");
document.write("3-slopeformula2 \n");
document.write("4-quadraticformula \n");
document.write("5-laplaceformula \n");
      
let choice=parseInt(prompt('Enter your choice '));

switch(choice) {
  case 1:
    const n = 5;
    const x = 4
    var i = 1;
    while(i<=n){
      const a  = Math.sin(n*x)
      const number = 1/n*a;
      console.log(number);
      i++;
    }
    break;

  case 2:
    let x1=parseInt(prompt("Enter first xcoordinate:"));
    let x2=parseInt(prompt("Enter the second xcoordinate:"));
    let y1=parseInt(prompt("Enter the first ycoordinate:"));
    let y2=parseInt(prompt("Enter the second ycoordinate:"));

    let output = slope_formula(x1,x2,y1,y2);
    console.log('The slope of the given coordinates is ',output );
    break;

  case 3:
    let d = parseInt(prompt('Enter the first number '));
    let e = parseInt(prompt('Enter the second number '));

    let result = square(d, e);
    console.log('The slope of the given coordinates is ', result );
    break;

  case 4:
    let a = parseInt(prompt('Enter the first number '));
    let b = parseInt(prompt('Enter the second number '));
    let c = parseInt(prompt('Enter the third number '));
    quadEquationSolver(a, b, c)
    break;
  
  case 5:
    var f=function(x,y){
      return x*x+2*x*y;
    }
    var partialDerivative=function(f,x,y,flag){
      var h = dx=dy=0.0001;
      var dz;
      switch (flag){
      case "x":
      dz=f(x+h,y)-f(x,y);
      
      return dz/dx;
      break;
      case "y":
      dz=f(x,y+h)-f(x,y);
      return dz/dy;
      break;
    }
    }
    alert(partialDerivative(f,2,1,"x"));
    //for sine and cosine functions
    var g=function(x,y){
      return Math.sin(x*y)-x*y*y;
    }
    var partialDerivative=function(f,x,y,flag){
      var h = dx=dy=0.0001;
      var dz;
      switch (flag){
      case "x":
      dz=g(x+h,y)-g(x,y);
      
      return dz/dx;
      break;
      case "y":
      dz=g(x,y+h)-g(x,y);
      return dz/dy;
      break;
    }
    }
    
    alert(partialDerivative(g,Math.PI,1,"x"));
    
    //for power functions
    var p=function(x,y){
      return Math.pow(6*x-y,4);
    }
    var partialDerivative=function(f,x,y,flag){
      var h = dx=dy=0.0001;
      var dz;
      switch (flag){
      case "x":
      dz=p(x+h,y)-p(x,y);
      
      return dz/dx;
      break;
      case "y":
      dz=p(x,y+h)-p(x,y);
      return dz/dy;
      break;
    }
    }
    
    alert(partialDerivative(p,1,4,"x"));
    break;
  
  default:
    hello("Choice...not found");
}
//Task01
//Quadratic Formula
function quadEquationSolver() {
    var a = 2.3, b = 4, c = 5.6;
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
//Task02
//Slopeformula (a-b)^2
var a = prompt("Enter a:");
var b = prompt("Enter b:");

a = parseInt(a);
b = parseInt(b);

function slopeformula(a, b) {
  return Math.pow(a - b, 2);
}

console.log(slopeformula(a, b));
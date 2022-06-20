const prompt=require("prompt-sync")({fake_val:"OPTIONAL CONFIG VALUES HERE",});
 let x1=parseInt(prompt("Enter first xcoordinate:"));
 let x2=parseInt(prompt("Enter the second xcoordinate:"));
 let y1=parseInt(prompt("Enter the first ycoordinate:"));
let y2=parseInt(prompt("Enter the second ycoordinate:"));
let gradient= (y2-y1)/(x2-x1)
console.log('The slope is ',gradient );

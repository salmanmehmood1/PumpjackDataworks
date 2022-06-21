const prompt=require("prompt-sync")({fake_val:"OPTIONAL CONFIG VALUES HERE",});
 let x1=parseInt(prompt("Enter first xcoordinate:"));
 let x2=parseInt(prompt("Enter the second xcoordinate:"));
 let y1=parseInt(prompt("Enter the first ycoordinate:"));
let y2=parseInt(prompt("Enter the second ycoordinate:"));
function slope_formula(x11,x22,y11,y22){
    let gradient= (y2-y1)/(x2-x1);
    return gradient;
}
let output=slope_formula(x1,x2,y1,y2);
console.log('The slope of the given coordinates is ',output );

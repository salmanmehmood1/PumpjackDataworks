//task01
//laplace in js
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


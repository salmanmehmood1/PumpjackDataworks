//task 2
//slope formula (a+b)^2
const prompt = require("prompt-sync")({
  fake_val: "OPTIONAL CONFIG VALUES HERE",
});

let a = parseInt(prompt("ENTER FIRST VALUE : "));
let b = parseInt(prompt("ENTER SECOND VALUE : "));

function square(d, e) {
  let c = Math.pow(d + e, 2);

  return c;
}

console.log(square(a, b));

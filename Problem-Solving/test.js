// let arr = [1, 2, 3, 4, 5];

// // arr = arr.map(index => index * 5);

// // arr = arr.filter(index => index !== '5')

// // arr = arr.reduce((acc, curr) => {
// //     return acc + curr
// // }, 0);

// arr.forEach((element, index) =>  {return arr[index] =  element * 2});

// console.log(arr)

let digit = "12345";

arr = digit.split("")

let sum = 0;

sum = arr.reduce((acc, curr) => {return acc + Number(curr)},0)

console.log(sum)

let digit2 = 12345;

let sum2 = 0;

while (digit2 !== 0){
    sum2 += digit2 % 10;
    digit2 = Math.floor(digit2/10);
}

console.log(sum2)


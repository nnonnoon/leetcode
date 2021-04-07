let nums = [2,5,1,3,4,7];
let n = 3;

let x = [];
let y = [];

for(let i = 0 ; i < n ; i++ ){
    x.push(nums[i])
}

for(let j = n ; j < nums.length ; j++){
    y.push(nums[j])  
}

// console.log(x)
// console.log(y)

let sum = [];

for(let k = 0 ; k < x.length ; k++){
    sum.push(x[k])
    sum.push(y[k])
}

// console.log(sum)

let n = 5;
let start = 0;

let res = [];
let tmp = 0;

for(let i = 0 ; i < n ; i++){
    tmp = start + 2 * i
    res.push(tmp);
}

res = res.reduce((acc, curr) =>  acc ^ curr);

console.log(res);
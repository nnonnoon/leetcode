let nums = [3,1,2,10,1]
let res = [];
let tmp = 0;

for(let i = 0 ; i < nums.length ; i++){
    tmp = tmp + nums[i];
    res.push(tmp)
}

console.log(res)

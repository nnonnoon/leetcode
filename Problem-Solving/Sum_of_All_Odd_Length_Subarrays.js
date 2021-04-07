let arr = [10,11,12]

let sum = 0;

for(let i = 0 ; i < arr.length ; i++){
    for(let j = i ; j < arr.length ; j++){
        if((j-i) % 2 === 0 ){
           for(let k = i ; k < j + 1 ; k++){
                sum += arr[k];
           }
        }
    }
}

console.log(sum)
let n = String(1500);
let arr = n.split("")

let tmp = "";
let digit = arr.length - 1;

for(let i = 0 ; i < arr.length ; i++){
    if(arr[i] === "0"){
        if(i === arr.length - 1){
            tmp = tmp.slice(0, -1)
        }
        digit -= 1;
        continue
    }
    
    tmp = tmp + arr[i] * (10 ** digit);

    if(i !== arr.length - 1){
        tmp = tmp + "+"
    }

    digit -= 1;
}

console.log(tmp)
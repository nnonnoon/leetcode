// let N = 67185;

// // const result = N.toString(2)

// // console.log(result)

// let bin = 0;
// let rem, i = 1 ;

// while(N !== 0){
//     rem = N % 2
//     N = Math.floor(N / 2)

//     bin += rem * i;
//     i = i* 10;
// }

// console.log(bin)


// bin = String(bin)
// new_bin = bin.split("")

// let res = "";
// let final = 0;
// let j = 0;

// for(let i = new_bin.length - 1  ; i > -1 ; i--){
//     if(new_bin[i] === "1"){
//         res = "0" + res;
//         final += 0 * (2**j);
//     }
//     else if (new_bin[i] === "0"){
//         res = "1" + res;
//         final += 1 * (2**j);
//     }
//     j += 1;
// }

// console.log(res)
// console.log(final)

    
var bitwiseComplement = function(N) {    
    let powerOf2 = 2;
    
    while(1){
        if (powerOf2 > N) {
            return powerOf2 - 1 - N;
        }
        powerOf2 *= 2        
    }
};

let res = bitwiseComplement(5)
console.log(res)

console.log(parseInt("2px"))





let s = "leetcode"

let keep_str = [];
let keep_int = [];
let new_str = "";

for(let i = 0 ; i < s.length ; i++){
    if(s.charCodeAt(i) >= 97 &&  s.charCodeAt(i) <= 122 ){
        keep_str.push(s[i]);
    }
    if(s.charCodeAt(i) >= 48 &&  s.charCodeAt(i) <= 57 ){
        keep_int.push(s[i]);
    }
}

if(keep_str.length === keep_int.length){
    for(let i = 0 ; i < keep_int.length ; i++){
        new_str = new_str.concat(keep_str[i], keep_int[i]);
    }

}

else if(keep_str.length - 1 === keep_int.length){
    for(let i = 0 ; i < keep_int.length ; i++){
        new_str = new_str.concat(keep_str[i], keep_int[i]);
    }
    new_str = new_str.concat(keep_str[keep_str.length - 1]);
}

else if(keep_int.length - 1 === keep_str.length){
    for(let i = 0 ; i < keep_str.length ; i++){
        new_str = new_str.concat(keep_str[i], keep_int[i]);
    }
    new_str = new_str.concat(keep_str[keep_int.length - 1]);
}else{
    new_str = '""';
}

console.log(new_str)


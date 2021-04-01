let arr = [1, 2, 3, 4, 5];

// arr = arr.map(index => index * 5);

// arr = arr.filter(index => index !== '5')

// arr = arr.reduce((acc, curr) => {
//     return acc + curr
// }, 0);

arr.forEach((element, index) =>  {return arr[index] =  element * 2});

console.log(arr)


let salary = [4000,3000,1000,2000];

let max_salary = Math.max(...salary);
let min_salary = Math.min(...salary);


let new_array = salary.filter(element => element !== max_salary && element !== min_salary);
let avg_salary = new_array.reduce((acc, curr) => (acc + curr));

console.log(avg_salary/new_array.length)



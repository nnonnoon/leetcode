class Solution {
    public String intToRoman(int num) {
        
        // num = 1994;
        
        String [] Symbol = {"I", "IV", "V", "IX","X", "XL","L", "XC","C", "CD","D", "CM","M"};
        int [] Value = {1, 4, 5, 9, 10, 40, 50, 90,100, 400, 500, 900, 1000};
        
        String result = "";
        
        int count = Value.length - 1;
        
        while(num > 0){
            int tmp = num / Value[count];  // หารเอาผล
            num = num % Value[count];   // หารเอาเศษ
            
            while(tmp > 0){
                result = result.concat(Symbol[count]);
                tmp = tmp - 1;
            }
            count = count - 1;
        }
        
        return result;
        
    }
}
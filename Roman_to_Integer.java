class Solution {
    public int romanToInt(String s) {
        
        StringBuffer new_s = new StringBuffer(s);
        
        HashMap<String, Integer> Roman_to_Integer = new HashMap<String, Integer>();
        
        Roman_to_Integer.put("I", 1);
        Roman_to_Integer.put("IV", 4);
        Roman_to_Integer.put("V", 5);
        Roman_to_Integer.put("IX", 9); 
        Roman_to_Integer.put("X", 10);
        Roman_to_Integer.put("XL", 40);
        Roman_to_Integer.put("L", 50);
        Roman_to_Integer.put("XC", 90);
        Roman_to_Integer.put("C", 100);
        Roman_to_Integer.put("CD", 400);
        Roman_to_Integer.put("D", 500);
        Roman_to_Integer.put("CM", 900);
        Roman_to_Integer.put("M", 1000);
        
        int result = 0;
        int count = 0 ;
        
        while(new_s.length() > 0){
            
            try {
              String check_symbol_1 = new_s.substring(0, 1);
              String check_symbol_2 = new_s.substring(0, 2);
                
                if(Roman_to_Integer.containsKey(check_symbol_2)){
                    result = result + Roman_to_Integer.get(check_symbol_2);
                    new_s.delete(0, 2);
                    continue;
                }
            
                if(Roman_to_Integer.containsKey(check_symbol_1)){
                    result = result + Roman_to_Integer.get(check_symbol_1);
                    new_s.delete(0, 1);
                }
            }
            catch(Exception e) {
                String check_symbol_1 = new_s.substring(0, 1);
                if(Roman_to_Integer.containsKey(check_symbol_1)){
                    result = result + Roman_to_Integer.get(check_symbol_1);
                    new_s.delete(0, 1);
                }

            }   
        }
        return result;
        
    }
}
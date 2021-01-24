class Solution {
    public boolean isPalindrome(int x) {
        
        String tmp = String.valueOf(x);
        
        String first = "";
        String second = "";
        
        if(tmp.length() == 1){
            return true;
        }
        
        for(int i = 0 ; i < tmp.length()/2 ; i++ ){
             char c =  tmp.charAt(i);
             first = first.concat(String.valueOf(c));
        }
        
        int range;
        if(tmp.length()%2 == 1){
            range = tmp.length()/2;
        }else{
            range = (tmp.length()/2) - 1;
        }
        
        for(int i = tmp.length()-1  ; i > range ; i-- ){
             char c =  tmp.charAt(i);
             second = second.concat(String.valueOf(c));
        }
        
        // System.out.print(first);
        // System.out.print("**");
        // System.out.print(second);
        // System.out.print(first.equals(second));
        
        if(first.equals(second)) {
            return true;
        }

        return false;
    }
}
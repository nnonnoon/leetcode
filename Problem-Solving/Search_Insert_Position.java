class Solution {
    public int searchInsert(int[] nums, int target) {
        
        // target = 0;
        
        int result = 0;
        
        for(int i = 0 ; i < nums.length ; i++){
            
            if(target > nums[nums.length - 1 ] ){
                result = nums.length;
                break;
            }
            
            else if(target < nums[0]){
                result = 0;
                break;
            }
            
            else if(nums[i] == target){
                result = i;
                break;
            }else{
                if(target < nums[i]){
                    result = i;
                    break;
                }
            }
            
        }
        return result;
    }
}
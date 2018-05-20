class Solution {
    public int[] twoSum(int[] nums, int target) {
        int [] response = new int [2];
        Map<Integer, Integer> arrayMap = new HashMap<>();
        for(Integer index = 0; index < nums.length; index += 1) {
            int needleValue = target - nums[index];
            if(arrayMap.containsKey(needleValue)) {
                int needleKey = arrayMap.get(needleValue);
                if(needleKey != index) {
                    response[0] = needleKey;
                    response[1] = index;
                    return response;
                }
           }
            arrayMap.put(nums[index], index);
        }
        return response;
    }
}

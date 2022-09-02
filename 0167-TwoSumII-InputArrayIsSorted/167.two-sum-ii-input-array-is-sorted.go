/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 *
 * https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
 *
 * algorithms
 * Easy (54.71%)
 * Likes:    1957
 * Dislikes: 650
 * Total Accepted:    463.4K
 * Total Submissions: 846.5K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers that is already sorted in ascending order, find
 * two numbers such that they add up to a specific target number.
 * 
 * The function twoSum should return indices of the two numbers such that they
 * add up to the target, where index1 must be less than index2.
 * 
 * Note:
 * 
 * 
 * Your returned answers (both index1 and index2) are not zero-based.
 * You may assume that each input would have exactly one solution and you may
 * not use the same element twice.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: numbers = [2,7,11,15], target = 9
 * Output: [1,2]
 * Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: numbers = [2,3,4], target = 6
 * Output: [1,3]
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: numbers = [-1,0], target = -1
 * Output: [1,2]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 2 <= nums.length <= 3 * 10^4
 * -1000 <= nums[i] <= 1000
 * nums is sorted in increasing order.
 * -1000 <= target <= 1000
 * 
 * 
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			for right-1 > 0 && numbers[right] == numbers[right-1] {
				right -= 1
			}
			right -= 1
		} else if sum < target {
			for left+1 < len(numbers) && numbers[left] == numbers[left+1] {
				left += 1
			}
			left += 1
		} else {
			return []int{left+1, right+1}
		}
	}

	return []int{}
}
// @lc code=end


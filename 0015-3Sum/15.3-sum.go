/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (27.25%)
 * Likes:    8242
 * Dislikes: 892
 * Total Accepted:    1.1M
 * Total Submissions: 3.9M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 * 
 * Notice that the solution set must not contain duplicate triplets.
 * 
 * 
 * Example 1:
 * Input: nums = [-1,0,1,2,-1,-4]
 * Output: [[-1,-1,2],[-1,0,1]]
 * Example 2:
 * Input: nums = []
 * Output: []
 * Example 3:
 * Input: nums = [0]
 * Output: []
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 * 
 * 
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	resList := make([][]int, 0, len(nums)/2)
	for idx := 0; idx < len(nums); idx += 1 {
		for _, existIdx := range twoSum(nums[idx+1:], -nums[idx]) {
			resList = append(resList, []int{nums[idx], existIdx[0], existIdx[1]})
		}
		for idx+1 < len(nums) && nums[idx] == nums[idx+1] {
			idx += 1
		}
	}
	return resList
}

func twoSum(nums []int, target int) [][]int{
	left, right := 0, len(nums)-1
	resList := make([][]int, 0, len(nums)/2)

	for left < right {
		sum := nums[left] + nums[right]
		if sum > target {
			for right-1 > 0 && nums[right] == nums[right-1] {
				right -= 1
			}
			right -= 1
		} else if sum < target {
			for left+1 < len(nums) && nums[left] == nums[left+1] {
				left += 1
			}
			left += 1
		} else {
			resList = append(resList, []int{nums[left], nums[right]})

			for left+1 < len(nums) && nums[left] == nums[left+1] {
				left += 1
			}
			left += 1

			for right-1 > 0 && nums[right] == nums[right-1] {
				right -= 1
			}
			right -= 1
		}

	}

	return resList
}
// @lc code=end


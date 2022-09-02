/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (34.06%)
 * Likes:    2445
 * Dislikes: 355
 * Total Accepted:    359.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers and an integer target, are there elements
 * a, b, c, and d in nums such that a + b + c + d = target? Find all unique
 * quadruplets in the array which gives the sum of target.
 * 
 * Notice that the solution set must not contain duplicate quadruplets.
 * 
 * 
 * Example 1:
 * Input: nums = [1,0,-1,0,-2,2], target = 0
 * Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 * Example 2:
 * Input: nums = [], target = 0
 * Output: []
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * 
 * 
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	resList := make([][]int, 0, len(nums)/2)
	for idx := 0; idx < len(nums); idx += 1 {
		for _, existIdx := range threeSum(nums[idx+1:], target-nums[idx]) {
			resList = append(resList, []int{nums[idx], existIdx[0], existIdx[1], existIdx[2]})
		}
		for idx+1 < len(nums) && nums[idx] == nums[idx+1] {
			idx += 1
		}
	}
	return resList
}

func threeSum(nums []int, target int) [][]int {
	resList := make([][]int, 0, len(nums)/2)
	for idx := 0; idx < len(nums); idx += 1 {
		for _, existIdx := range twoSum(nums[idx+1:], target-nums[idx]) {
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


/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
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

func twoSum(nums []int, target int) [][]int {
	resList := make([][]int, 0, len(nums)/2)

	for left, right := 0, len(nums)-1; left < right; {
		sum := nums[left] + nums[right]
		if sum > target {
			for right-1 > left && nums[right] == nums[right-1] {
				right -= 1
			}
			right -= 1
		} else if sum < target {
			for left+1 < right && nums[left] == nums[left+1] {
				left += 1
			}
			left += 1
		} else {
			resList = append(resList, []int{nums[left], nums[right]})

			for left+1 < right && nums[left] == nums[left+1] {
				left += 1
			}
			left += 1

			for right-1 > left && nums[right] == nums[right-1] {
				right -= 1
			}
			right -= 1
		}
	}

	return resList
}

// @lc code=end


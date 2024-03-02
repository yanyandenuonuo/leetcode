/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

// @lc code=start
func sortArrayByParityII(nums []int) []int {
	for leftIdx, rightIdx := 0, len(nums)-1; leftIdx < len(nums) && rightIdx >= 0; {
		if leftIdx&0x01 == 0x00 && nums[leftIdx]&0x01 == 0x00 {
			leftIdx += 2
			continue
		}

		if rightIdx&0x01 == 0x01 && nums[rightIdx]&0x01 == 0x01 {
			rightIdx -= 2
			continue
		}

		nums[leftIdx], nums[rightIdx] = nums[rightIdx], nums[leftIdx]

		leftIdx += 2
		rightIdx -= 2
	}

	return nums
}

// @lc code=end


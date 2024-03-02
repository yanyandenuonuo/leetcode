/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */

// @lc code=start
func sortArrayByParity(nums []int) []int {
	for leftIdx, rightIdx := 0, len(nums)-1; leftIdx < rightIdx; {
		if nums[leftIdx]&0x01 == 0x00 {
			leftIdx += 1
			continue
		}

		if nums[rightIdx]&0x01 == 0x01 {
			rightIdx -= 1
			continue
		}

		nums[leftIdx], nums[rightIdx] = nums[rightIdx], nums[leftIdx]

		leftIdx += 1
		rightIdx -= 1
	}

	return nums
}

// @lc code=end


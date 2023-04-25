/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	numLength := len(nums)
	for idx := 0; idx < numLength; {
		if nums[idx] != val {
			idx += 1
			continue
		}

		if nums[numLength-1] == val {
			numLength -= 1
			continue
		}

		nums[idx] = nums[numLength-1]
		idx += 1
		numLength -= 1
	}
	return numLength
}

// @lc code=end


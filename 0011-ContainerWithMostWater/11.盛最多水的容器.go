/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	// solution: 双指针
	maxRes, leftIdx, rightIdx := 0, 0, len(height)-1

	for leftIdx < rightIdx {
		res := 0
		if height[leftIdx] < height[rightIdx] {
			res = (rightIdx - leftIdx) * height[leftIdx]
			leftIdx += 1
		} else {
			res = (rightIdx - leftIdx) * height[rightIdx]
			rightIdx -= 1
		}

		if res > maxRes {
			maxRes = res
		}
	}

	return maxRes
}

// @lc code=end


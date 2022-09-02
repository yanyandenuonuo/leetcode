/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	leftIdx, rightIdx := 0, len(nums)-1
	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2

		preVal := 0
		if midIdx-1 < 0 {
			preVal = nums[midIdx] - 1
		} else {
			preVal = nums[midIdx-1]
		}

		nextVal := 0
		if midIdx+1 < len(nums) {
			nextVal = nums[midIdx+1]
		} else {
			nextVal = nums[midIdx] - 1
		}

		if preVal < nums[midIdx] && nums[midIdx] > nextVal {
			return midIdx
		}
		if nums[midIdx] < nums[midIdx+1] {
			leftIdx = midIdx + 1
		} else {
			rightIdx = midIdx - 1
		}
	}

	return -1
}

// @lc code=end


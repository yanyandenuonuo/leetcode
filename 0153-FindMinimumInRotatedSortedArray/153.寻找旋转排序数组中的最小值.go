/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	minIdx, highIdx := 0, len(nums)-1
	for minIdx < highIdx {
		midIdx := minIdx + (highIdx-minIdx)/2

		// 旋转点在右侧
		if nums[midIdx] > nums[highIdx] {
			minIdx = midIdx + 1
		} else {
			highIdx = midIdx
		}
	}
	return nums[minIdx]
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	res := len(nums)
	for leftIdx, rightIdx := 0, len(nums)-1; leftIdx <= rightIdx; {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		if nums[midIdx] > target {
			rightIdx = midIdx - 1
			res = midIdx
		} else if nums[midIdx] < target {
			leftIdx = midIdx + 1
		} else {
			return midIdx
		}
	}
	return res
}

// @lc code=end


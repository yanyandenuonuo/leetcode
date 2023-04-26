/*
 * @lc app=leetcode.cn id=845 lang=golang
 *
 * [845] 数组中的最长山脉
 */

// @lc code=start
func longestMountain(arr []int) int {
	// solution1: 枚举左山脚和右山脚，通过双指针移动算出符合条件的山脉
	// solution2: 枚举山顶，分别求出左边的山脉长度和右边的山脉长度
	leftMountain := make([]int, len(arr))
	leftMountain[0] = 0
	for idx := 1; idx < len(leftMountain); idx += 1 {
		if arr[idx] > arr[idx-1] {
			leftMountain[idx] = leftMountain[idx-1] + 1
		}
	}

	rightMountain := make([]int, len(arr))
	rightMountain[len(arr)-1] = 0
	for idx := len(arr) - 2; idx >= 0; idx -= 1 {
		if arr[idx] > arr[idx+1] {
			rightMountain[idx] = rightMountain[idx+1] + 1
		}
	}

	res := 0
	for idx := 1; idx < len(arr)-1; idx += 1 {
		if leftMountain[idx] > 0 && rightMountain[idx] > 0 && leftMountain[idx]+rightMountain[idx]+1 > res {
			res = leftMountain[idx] + rightMountain[idx] + 1
		}
	}
	return res
}

// @lc code=end


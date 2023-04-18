/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */

// @lc code=start
func validMountainArray(arr []int) bool {
	idx := 0

	// check increase
	for ; idx < len(arr)-1 && arr[idx] < arr[idx+1]; idx += 1 {

	}

	// peak should not equal first or end
	if idx == 0 || idx == len(arr)-1 {
		return false
	}

	// check decrease
	for ; idx < len(arr)-1 && arr[idx] > arr[idx+1]; idx += 1 {

	}

	return idx == len(arr)-1
}

// @lc code=end


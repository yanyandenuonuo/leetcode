/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	leftIdx, rightIdx := 0, len(arr)-1
	for leftIdx <= rightIdx {
		midIdx := leftIdx + (rightIdx-leftIdx)/2

		if (midIdx-1 < 0 || arr[midIdx-1] < arr[midIdx]) && (midIdx+1 > len(arr)-1 || arr[midIdx] > arr[midIdx+1]) {
			return midIdx
		}
		if arr[midIdx] < arr[midIdx+1] {
			leftIdx = midIdx + 1
		} else {
			rightIdx = midIdx - 1
		}
	}
	return -1
}

// @lc code=end


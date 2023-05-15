/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] 寻找比目标字母大的最小字母
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	// solution1: 遍历数组查找
	// solution2: 二分查找
	leftIdx := 0
	for rightIdx := len(letters) - 1; leftIdx <= rightIdx; {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		if letters[midIdx] > target {
			rightIdx = midIdx - 1
		} else {
			leftIdx = midIdx + 1
		}
	}

	if leftIdx == len(letters) {
		return letters[0]
	}

	return letters[leftIdx]
}

// @lc code=end


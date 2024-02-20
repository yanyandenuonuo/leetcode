/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] 字符的最短距离
 */

// @lc code=start
func shortestToChar(s string, c byte) []int {
	// solution: 左右遍历，距离左边c的距离，距离右边c的距离，取较小值
	offset := make([]int, len(s))

	targetIdx := -1 << 16
	for idx := 0; idx < len(s); idx += 1 {
		if s[idx] == c {
			targetIdx = idx
			offset[idx] = 0
			continue
		}

		offset[idx] = idx - targetIdx
	}

	targetIdx = 1 << 16
	for idx := len(s) - 1; idx >= 0; idx -= 1 {
		if s[idx] == c {
			targetIdx = idx
			offset[idx] = 0
			continue
		}

		if offset[idx] > targetIdx-idx {
			offset[idx] = targetIdx - idx
		}
	}

	return offset
}

// @lc code=end


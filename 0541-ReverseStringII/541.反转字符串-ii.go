/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	// solution: 模拟
	charBytes := []byte(s)

	for idx := 0; idx < len(s); idx += 2 * k {
		rightIdx := idx + k - 1
		if rightIdx >= len(s) {
			rightIdx = len(s) - 1
		}

		for leftIdx := idx; leftIdx < rightIdx; {
			charBytes[leftIdx], charBytes[rightIdx] = charBytes[rightIdx], charBytes[leftIdx]
			leftIdx += 1
			rightIdx -= 1
		}
	}

	return string(charBytes)
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=696 lang=golang
 *
 * [696] 计数二进制子串
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	// solution: 计数，然后统计相邻子串的较小值计入结果
	counters := make([]int, 0)
	for idx := 0; idx < len(s); {
		rightIdx := idx + 1
		for rightIdx < len(s) && s[idx] == s[rightIdx] {
			rightIdx += 1
		}

		counters = append(counters, rightIdx-idx)
		idx = rightIdx
	}

	counter := 0

	for idx := 1; idx < len(counters); idx += 1 {
		if counters[idx-1] < counters[idx] {
			counter += counters[idx-1]
		} else {
			counter += counters[idx]
		}
	}

	return counter
}

// @lc code=end


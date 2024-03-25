/*
 * @lc app=leetcode.cn id=1529 lang=golang
 *
 * [1529] 最少的后缀翻转次数
 */

// @lc code=start
func minFlips(target string) int {
	// solution: 不同区段计数
	counter := 0
	for currentBit, idx := byte('0'), 0; idx < len(target); idx += 1 {
		if currentBit == target[idx] {
			continue
		}

		currentBit = target[idx]

		counter += 1
	}

	return counter
}

// @lc code=end


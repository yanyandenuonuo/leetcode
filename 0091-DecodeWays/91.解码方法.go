/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	// solution: 动态规划
	pre1Count, preCount := 0, 1
	preChar := '0'
	for _, currentChar := range s {
		if currentChar == '0' {
			// 只能与前一个字符组成2位：10或者20
			if preChar != '1' && preChar != '2' {
				return 0
			}

			// 这里实际应该为
			// 	preCount = pre1Count
			// 	count = preCount
			// 	pre1Count, preCount = preCount, count
			pre1Count, preCount = pre1Count, pre1Count
			preChar = currentChar
			continue
		}

		// 单独1位
		count := preCount

		// 与前一个字符结合组成2位
		if preChar == '1' || (preChar == '2' && '1' <= currentChar && currentChar <= '6') {
			count += pre1Count
		}

		pre1Count, preCount = preCount, count
		preChar = currentChar
	}

	return preCount
}

// @lc code=end


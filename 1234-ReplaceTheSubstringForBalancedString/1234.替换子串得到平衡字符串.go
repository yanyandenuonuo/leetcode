/*
 * @lc app=leetcode.cn id=1234 lang=golang
 *
 * [1234] 替换子串得到平衡字符串
 */

// @lc code=start
func balancedString(s string) int {
	// solution: 滑动窗口
	maxCount := len(s) / 4
	charCountMap := make(map[byte]int, 4)

	for idx := range s {
		charCountMap[s[idx]] += 1
	}

	if isValid(charCountMap, maxCount) {
		// 初始已经满足条件
		return 0
	}

	minSubLength := len(s)

	rightIdx := 0
	for leftIdx := 0; leftIdx < len(s); leftIdx += 1 {
		for ; !isValid(charCountMap, maxCount) && rightIdx < len(s); rightIdx += 1 {
			charCountMap[s[rightIdx]] -= 1
		}

		if !isValid(charCountMap, maxCount) {
			// 遍历到最后仍未满足条件
			break
		}

		if rightIdx-leftIdx < minSubLength {
			minSubLength = rightIdx - leftIdx
		}

		charCountMap[s[leftIdx]] += 1
	}
	return minSubLength
}

func isValid(charCountMap map[byte]int, maxCount int) bool {
	if charCountMap['Q'] > maxCount || charCountMap['W'] > maxCount || charCountMap['E'] > maxCount || charCountMap['R'] > maxCount {
		return false
	}

	return true
}

// @lc code=end


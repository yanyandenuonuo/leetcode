/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	// solution1: get real string, then compare
	// solution2: compare from the string end
	for sIdx, tIdx := len(s)-1, len(t)-1; sIdx >= 0 || tIdx >= 0; {
		backspaceCount := 0
		for sIdx >= 0 && (s[sIdx] == '#' || backspaceCount > 0) {
			if s[sIdx] != '#' {
				backspaceCount -= 1
			} else {
				backspaceCount += 1
			}
			sIdx -= 1
		}

		backspaceCount = 0
		for tIdx >= 0 && (t[tIdx] == '#' || backspaceCount > 0) {
			if t[tIdx] != '#' {
				backspaceCount -= 1
			} else {
				backspaceCount += 1
			}
			tIdx -= 1
		}

		if sIdx >= 0 && tIdx >= 0 && s[sIdx] == t[tIdx] {
			sIdx -= 1
			tIdx -= 1
		} else if sIdx >= 0 || tIdx >= 0 {
			return false
		}
	}

	return true
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	sRunes, tRunes := []rune(s), []rune(t)
	sIdx, tIdx := len(sRunes)-1, len(tRunes)-1
	for sIdx >= 0 || tIdx >= 0 {
		backspaceCount := 0
		for sIdx >= 0 && (sRunes[sIdx] == '#' || backspaceCount > 0) {
			if sRunes[sIdx] != '#' {
				backspaceCount -= 1
			} else {
				backspaceCount += 1
			}
			sIdx -= 1
		}

		backspaceCount = 0
		for tIdx >= 0 && (tRunes[tIdx] == '#' || backspaceCount > 0) {
			if tRunes[tIdx] != '#' {
				backspaceCount -= 1
			} else {
				backspaceCount += 1
			}
			tIdx -= 1
		}

		if sIdx >= 0 && tIdx >= 0 && sRunes[sIdx] == tRunes[tIdx] {
			// fmt.Printf("sIdx:%v, tIdx:%v, sRune:%v, tRune:%v \n", sIdx, tIdx, string(sRunes[sIdx]), string(tRunes[tIdx]))

			sIdx -= 1
			tIdx -= 1
		} else {
			if sIdx >= 0 || tIdx >= 0 {
				return false
			}
		}
	}

	return true
}

// @lc code=end


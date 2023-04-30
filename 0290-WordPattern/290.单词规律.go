/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	// solution: hash映射
	wordList := strings.Split(s, " ")
	if len(pattern) != len(wordList) {
		return false
	}
	wordMap := make(map[byte]string, len(pattern))
	strMap := make(map[string]byte, len(pattern))
	for idx := range wordList {
		if len(wordMap[pattern[idx]]) == 0 {
			if strMap[wordList[idx]] > 0 {
				return false
			}

			wordMap[pattern[idx]] = wordList[idx]
			strMap[wordList[idx]] = pattern[idx]
			continue
		}
		if wordMap[pattern[idx]] != wordList[idx] {
			return false
		}
	}
	return true
}

// @lc code=end


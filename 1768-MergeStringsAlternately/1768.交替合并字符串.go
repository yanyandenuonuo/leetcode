/*
 * @lc app=leetcode.cn id=1768 lang=golang
 *
 * [1768] 交替合并字符串
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	// solution: 双指针
	word1Idx, word2Idx := 0, 0
	res := new(strings.Builder)
	for word1Idx < len(word1) && word2Idx < len(word2) {
		if word1Idx <= word2Idx {
			res.WriteByte(word1[word1Idx])
			word1Idx += 1
		} else {
			res.WriteByte(word2[word2Idx])
			word2Idx += 1
		}
	}

	if word1Idx < len(word1) {
		res.WriteString(word1[word1Idx:])
	}

	if word2Idx < len(word2) {
		res.WriteString(word2[word2Idx:])
	}

	return res.String()
}

// @lc code=end


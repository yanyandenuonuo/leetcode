/*
 * @lc app=leetcode.cn id=884 lang=golang
 *
 * [884] 两句话中的不常见单词
 */

// @lc code=start
func uncommonFromSentences(s1 string, s2 string) []string {
	// Solution1: 位标记
	// 			  在s1中出现1次置为0b01，在s1中出现置为0b0100
	// 			  在s2中出现1次置为0b10，在s1中出现置为0b1000

	wordMap := make(map[string]uint8, 0)

	for _, s1Word := range strings.Split(s1, " ") {
		if _, isExist := wordMap[s1Word]; !isExist {
			wordMap[s1Word] |= 0b01
		} else {
			wordMap[s1Word] |= 0b100
		}
	}

	for _, s2Word := range strings.Split(s2, " ") {
		if _, isExist := wordMap[s2Word]; !isExist {
			wordMap[s2Word] |= 0b10
		} else {
			wordMap[s2Word] |= 0b1000
		}
	}

	resList := make([]string, 0, len(wordMap))

	for word, val := range wordMap {
		if val == 0x01 || val == 0x02 {
			resList = append(resList, word)
		}
	}

	return resList
}

// @lc code=end


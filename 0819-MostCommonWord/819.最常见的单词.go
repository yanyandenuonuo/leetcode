/*
 * @lc app=leetcode.cn id=819 lang=golang
 *
 * [819] 最常见的单词
 */

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
	// solution: hash
	wordCounter := make(map[string]int, 0)
	for idx := 0; idx < len(paragraph); {
		if !(paragraph[idx] >= 'a' && paragraph[idx] <= 'z') && !(paragraph[idx] >= 'A' && paragraph[idx] <= 'Z') {
			idx += 1
			continue
		}
		rightIdx := idx + 1
		for ; rightIdx < len(paragraph); rightIdx += 1 {
			if !(paragraph[rightIdx] >= 'a' && paragraph[rightIdx] <= 'z') && !(paragraph[rightIdx] >= 'A' && paragraph[rightIdx] <= 'Z') {
				break
			}
		}

		word := paragraph[idx:rightIdx]
		idx = rightIdx + 1

		if len(word) == 0 {
			continue
		}

		wordCounter[strings.ToLower(word)] += 1
	}

	for _, word := range banned {
		delete(wordCounter, word)
	}

	maxFreqWord := ""

	for word, counter := range wordCounter {
		if counter > wordCounter[maxFreqWord] {
			maxFreqWord = word
		}
	}

	return maxFreqWord
}

// @lc code=end


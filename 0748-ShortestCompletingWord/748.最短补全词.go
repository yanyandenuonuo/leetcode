/*
 * @lc app=leetcode.cn id=748 lang=golang
 *
 * [748] 最短补全词
 */

// @lc code=start
func shortestCompletingWord(licensePlate string, words []string) string {
	// solution: 计数
	charCounter, charCount := getCharCounter(licensePlate)

	res := ""

	for _, word := range words {
		wordCounter, wordCount := getCharCounter(word)
		if wordCount < charCount {
			continue
		}

		isValid := true
		for idx := range wordCounter {
			if wordCounter[idx] < charCounter[idx] {
				isValid = false
				break
			}
		}

		if isValid && (res == "" || len(word) < len(res)) {
			res = word
		}
	}

	return res
}

func getCharCounter(word string) ([]byte, byte) {
	charCounter, charCount := make([]byte, 26), byte(0)
	for _, charRune := range word {
		if charRune >= 'a' && charRune <= 'z' {
			charCounter[charRune-'a'] += 1
			charCount += 1
		} else if charRune >= 'A' && charRune <= 'Z' {
			charCounter[charRune-'A'] += 1
			charCount += 1
		}
	}

	return charCounter, charCount
}

// @lc code=end


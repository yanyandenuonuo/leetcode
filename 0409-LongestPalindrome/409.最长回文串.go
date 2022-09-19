/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
func longestPalindrome(s string) int {
	total := 0
	runeMap := make(map[rune]int, len(s))
	for _, curRune := range s {
		if _, isExist := runeMap[curRune]; !isExist {
			runeMap[curRune] = 0
		}
		runeMap[curRune] += 1
		if runeMap[curRune]%2 == 0 {
			total += 2
		}
	}
	if total != len(s) {
		total += 1
	}
	return total
}

// @lc code=end


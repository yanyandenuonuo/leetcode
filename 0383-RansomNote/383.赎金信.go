/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	charMap := make(map[rune]int, len(magazine))
	for _, char := range magazine {
		charMap[char] += 1
	}
	for _, char := range ransomNote {
		charMap[char] -= 1
		if charMap[char] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end


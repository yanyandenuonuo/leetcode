/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	charMap := make(map[rune]int, len(s))
	for _, char := range s {
		charMap[char] += 1
	}
	for _, char := range t {
		charMap[char] -= 1
		if charMap[char] == 0 {
			delete(charMap, char)
		}
	}
	return len(charMap) == 0
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	runeMap := map[byte]bool{
		'a': true,
		'A': true,
		'e': true,
		'E': true,
		'i': true,
		'I': true,
		'o': true,
		'O': true,
		'u': true,
		'U': true,
	}
	sBytes := []byte(s)
	for leftIdx, rightIdx := 0, len(s)-1; leftIdx < rightIdx; {
		if runeMap[s[leftIdx]] && runeMap[s[rightIdx]] {
			sBytes[leftIdx], sBytes[rightIdx] = sBytes[rightIdx], sBytes[leftIdx]
			leftIdx += 1
			rightIdx -= 1
		}
		if !runeMap[s[leftIdx]] {
			leftIdx += 1
		}
		if !runeMap[s[rightIdx]] {
			rightIdx -= 1
		}
	}
	return string(sBytes)
}

// @lc code=end


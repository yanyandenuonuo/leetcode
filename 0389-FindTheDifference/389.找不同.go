/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	runeCount := [26]byte{}
	for idx := range s {
		runeCount[s[idx]-'a'] += 1
	}
	for idx := range t {
		if runeCount[t[idx]-'a'] == 0 {
			return t[idx]
		}
		runeCount[t[idx]-'a'] -= 1
	}
	for idx := range runeCount {
		if runeCount[idx] < 0 {
			return byte(idx + 'a')
		}
	}
	return ' '
}

// @lc code=end


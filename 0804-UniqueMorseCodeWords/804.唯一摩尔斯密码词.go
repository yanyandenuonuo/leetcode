/*
 * @lc app=leetcode.cn id=804 lang=golang
 *
 * [804] 唯一摩尔斯密码词
 */

// @lc code=start
func uniqueMorseRepresentations(words []string) int {
	// solution: hash
	charToMorse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	morseMap := make(map[string]bool, 0)
	for _, word := range words {
		morse := new(strings.Builder)
		for _, runeChar := range word {
			morse.WriteString(charToMorse[runeChar-'a'])
		}
		morseMap[morse.String()] = true
	}

	return len(morseMap)
}

// @lc code=end

